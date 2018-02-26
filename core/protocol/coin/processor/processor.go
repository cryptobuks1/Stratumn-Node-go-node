// Copyright © 2017-2018 Stratumn SAS
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:generate mockgen -package mockprocessor -destination mockprocessor/mockprocessor.go github.com/stratumn/alice/core/protocol/coin/processor Processor

package processor

import (
	"bytes"
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/stratumn/alice/core/protocol/coin/chain"
	"github.com/stratumn/alice/core/protocol/coin/coinutil"
	"github.com/stratumn/alice/core/protocol/coin/state"
	pb "github.com/stratumn/alice/pb/coin"

	logging "gx/ipfs/QmSpJByNKFX1sCsHBEp3R73FL4NF6FnQTEGyNAXHm2GS52/go-log"
	cid "gx/ipfs/QmcZfnkapfECQGcLZaf9B79NRg7cRa9EnZh4LSbkCzwNvY/go-cid"
)

var (
	// log is the logger for the processor.
	log = logging.Logger("coin.processor")

	// errReorgAborted is returned when the reorg is aborted.
	errReorgAborted = errors.New("reorg aborted")

	// ProvideTimeout is the timeout when letting the network we provide a resoruce.
	ProvideTimeout = time.Second * 10
)

// Processor is an interface for processing blocks using a given initial state.
type Processor interface {
	// Process applies the state changes from the block contents
	// and adds the block to the chain.
	Process(ctx context.Context, block *pb.Block, state state.State, ch chain.Chain) error
}

// ContentProvider is an interface used to let the network know we provide a resource.
// The resource is identified by a content ID.
type ContentProvider interface {
	Provide(ctx context.Context, key *cid.Cid, brdcst bool) error
}

type processor struct {
	// provider is used to let the network know we have a block once we added it to the local chain.
	provider ContentProvider
}

// NewProcessor creates a new processor.
func NewProcessor(provider ContentProvider) Processor {
	return &processor{provider: provider}
}

func (p *processor) Process(ctx context.Context, block *pb.Block, state state.State, ch chain.Chain) error {
	mh, err := coinutil.HashHeader(block.Header)
	if err != nil {
		return err
	}

	// Check block has already been processed.
	if _, err := ch.GetBlock(mh, block.BlockNumber()); err != nil && err != chain.ErrBlockNotFound {
		return err
	} else if err == nil {
		log.Event(context.Background(), "BlockAlreadyProcessed", logging.Metadata{"hash": mh.String(), "height": block.BlockNumber()})
		return nil
	}

	// Update chain.
	if err := ch.AddBlock(block); err != nil {
		return err
	}

	// Tell the network we have that block.
	if p.provider != nil {
		contentID, err := cid.Cast(mh)
		if err != nil {
			log.Event(ctx, "failCastHashToCID", logging.Metadata{"hash": mh.B58String()})
		} else {
			provideCtx, cancel := context.WithTimeout(ctx, ProvideTimeout)
			defer cancel()
			if err = p.provider.Provide(provideCtx, contentID, true); err != nil {
				log.Event(ctx, "failProvide", logging.Metadata{"cid": contentID.String(), "error": err.Error()})
			}
		}
	}

	head, err := ch.CurrentBlock()
	if err != nil && err != chain.ErrBlockNotFound {
		return err
	}

	// If the block is not higher than the current head,
	// do not update state, end processing.
	if head != nil && head.BlockNumber() >= block.BlockNumber() {
		return nil
	}

	// Set the new head.
	if err := ch.SetHead(block); err != nil {
		return err
	}

	// If the new head is on a different branch, reorganize state.
	if head != nil {
		hash, err := coinutil.HashHeader(head.Header)
		if err != nil {
			return err
		}
		if !bytes.Equal(hash, block.PreviousHash()) {
			if err := p.reorg(head, block, state, ch); err != nil {
				if err == errReorgAborted {
					log.Event(ctx, "ReorgAborted")

					return ch.SetHead(head)
				}

				return err
			}

			return nil
		}
	}

	// Update state.
	return state.ProcessBlock(block)
}

// Update the state to follow the new main branch.
func (p *processor) reorg(prevHead *pb.Block, newHead *pb.Block, st state.State, ch chain.Reader) error {
	var cursor *pb.Block

	backward, forward, err := chain.GetPath(ch, prevHead, newHead)
	if err != nil {
		return err
	}

	forward = append(forward, newHead)

	for _, b := range backward {
		if err = st.RollbackBlock(b); err != nil {
			return err
		}
		cursor = b
	}

	for _, b := range forward {
		if err = st.ProcessBlock(b); err != nil {
			// Block has been rejected by state, undo reorg.
			if err == state.ErrInvalidBlock {
				if cursor != nil {
					if err := p.reorg(cursor, prevHead, st, ch); err != nil {
						return err
					}
				}

				return errReorgAborted
			}

			return err
		}
		cursor = b
	}

	return nil
}
