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

package protector

import (
	"context"
	"sync"

	pb "github.com/stratumn/alice/pb/protector"

	"gx/ipfs/QmPUHzTLPZFYqv8WqcBTuMFYTgeom4uHHEaxzk7bd5GYZB/go-libp2p-transport"
	"gx/ipfs/QmWWQ2Txc2c6tqjsBpzg5Ar652cHPGNsQQp2SejkNmkUMb/go-multiaddr"
	"gx/ipfs/QmcJukH2sAFjY3HdBKq35WDzWoL3UUu2gt9wdfqZTUyM74/go-libp2p-peer"
	"gx/ipfs/Qmd3oYWVLCVWryDV6Pobv6whZcvDXAHqS3chemZ658y4a8/go-libp2p-interface-pnet"
	"gx/ipfs/QmdeiKhUy1TVGBaKxt7y1QmBDLBdisSrLJ1x58Eoj4PXUh/go-libp2p-peerstore"
)

// PrivateNetworkWithBootstrap implements the github.com/libp2p/go-libp2p-interface-pnet/ipnet.Protector interface.
// It protects a network by only allowing whitelisted peers to connect once the
// bootstrap phase is complete.
// During the bootstrap phase, it accepts all requests.
type PrivateNetworkWithBootstrap struct {
	privateNetwork Protector

	networkStateLock sync.RWMutex
	networkState     pb.NetworkState
}

// NewPrivateNetworkWithBootstrap creates a protector for private networks
// supporting an open bootstrapping phase.
// The protector accepts all connections during the bootstrap phase.
// Once the network state changes and ends the bootstrap phase, the protector
// starts rejecting every non-white-listed request.
func NewPrivateNetworkWithBootstrap(peerStore peerstore.Peerstore) Protector {
	p := PrivateNetworkWithBootstrap{
		privateNetwork: NewPrivateNetwork(peerStore),
		networkState:   pb.NetworkState_BOOTSTRAP,
	}

	// We initially allow all requests.
	ipnet.ForcePrivateNetwork = false

	return &p
}

// Protect accepts all connections until the bootstrap channel is notified.
// Then it switches to private network mode.
func (p *PrivateNetworkWithBootstrap) Protect(conn transport.Conn) (transport.Conn, error) {
	p.networkStateLock.RLock()
	bootstrapDone := p.networkState != pb.NetworkState_BOOTSTRAP
	p.networkStateLock.RUnlock()

	if !bootstrapDone {
		return conn, nil
	}

	return p.privateNetwork.Protect(conn)
}

// ListenForUpdates listens for network updates.
// This is a blocking call that should be made in a dedicated go routine.
// Closing the channel will stop the listener.
func (p *PrivateNetworkWithBootstrap) ListenForUpdates(updateChan <-chan NetworkUpdate) {
	p.privateNetwork.ListenForUpdates(updateChan)
}

// Fingerprint returns a hash of the participants list.
func (p *PrivateNetworkWithBootstrap) Fingerprint() []byte {
	return p.privateNetwork.Fingerprint()
}

// AllowedAddrs returns the list of whitelisted addresses.
func (p *PrivateNetworkWithBootstrap) AllowedAddrs(ctx context.Context) []multiaddr.Multiaddr {
	return p.privateNetwork.AllowedAddrs(ctx)
}

// AllowedPeers returns the list of whitelisted peers.
func (p *PrivateNetworkWithBootstrap) AllowedPeers(ctx context.Context) []peer.ID {
	return p.privateNetwork.AllowedPeers(ctx)
}

// SetNetworkState sets the network state. The protector adapts to the
// network state, so this method should be called when it changes.
func (p *PrivateNetworkWithBootstrap) SetNetworkState(_ context.Context, networkState pb.NetworkState) error {
	p.networkStateLock.Lock()
	defer p.networkStateLock.Unlock()

	p.networkState = networkState
	switch p.networkState {
	case pb.NetworkState_BOOTSTRAP:
		ipnet.ForcePrivateNetwork = false
	default:
		ipnet.ForcePrivateNetwork = true
	}

	return nil
}