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

package miner

import (
	"github.com/stratumn/alice/core/protocol/coin/chain"
	"github.com/stratumn/alice/core/protocol/coin/engine"
	"github.com/stratumn/alice/core/protocol/coin/processor"
	"github.com/stratumn/alice/core/protocol/coin/state"
	"github.com/stratumn/alice/core/protocol/coin/testutil"
	"github.com/stratumn/alice/core/protocol/coin/validator"
)

// MinerBuilder is a utility to create miners with custom mocks
// to simulate a wide variety of miner configuration.
type MinerBuilder struct {
	chain     chain.Chain
	engine    engine.Engine
	mempool   state.Mempool
	processor processor.Processor
	state     state.State
	validator validator.Validator
}

// NewMinerBuilder creates a MinerBuilder with a context and
// good default values to build a test Miner.
func NewMinerBuilder() *MinerBuilder {
	return &MinerBuilder{
		engine:    &testutil.DummyEngine{},
		mempool:   &testutil.InMemoryMempool{},
		validator: &testutil.DummyValidator{},
	}
}

// WithEngine configures the builder to use the given engine.
func (m *MinerBuilder) WithEngine(engine engine.Engine) *MinerBuilder {
	m.engine = engine
	return m
}

// WithMempool configures the builder to use the given mempool.
func (m *MinerBuilder) WithMempool(mempool state.Mempool) *MinerBuilder {
	m.mempool = mempool
	return m
}

// WithValidator configures the builder to use the given validator.
func (m *MinerBuilder) WithValidator(validator validator.Validator) *MinerBuilder {
	m.validator = validator
	return m
}

// Build builds the underlying miner and returns it.
// It's now ready to use in your tests.
func (m *MinerBuilder) Build() *Miner {
	return NewMiner(
		m.mempool,
		m.engine,
		m.state,
		m.chain,
		m.validator,
		m.processor,
	)
}
