// Copyright © 2017  Stratumn SAS
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

package metrics

import (
	"context"
	"testing"
	"time"

	metrics "github.com/armon/go-metrics"
	"github.com/pkg/errors"
	pb "github.com/stratumn/alice/grpc/metrics"

	p2pmetrics "gx/ipfs/QmQbh3Rb7KM37As3vkHYnEFnzkVXNCP8EYGtHz6g2fXk14/go-libp2p-metrics"
)

func testGRPCServer() grpcServer {
	mtrx := newMetrics(
		p2pmetrics.NewBandwidthCounter(),
		metrics.NewInmemSink(time.Second, time.Second),
	)
	return grpcServer{func() *Metrics { return mtrx }}
}

func testGRPCServerUnavailable() grpcServer {
	return grpcServer{func() *Metrics { return nil }}
}
func TestGRPCServer_Bandwidth(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	srv := testGRPCServer()

	srv.GetMetrics().LogRecvMessage(1000)

	req := &pb.BandwidthReq{}
	res, err := srv.Bandwidth(ctx, req)
	if err != nil {
		t.Fatalf("srv.Bandwidth(ctx, req): error: %s", err)
	}

	if got, want := res.TotalIn, uint64(1000); got != want {
		t.Errorf("res.TotalIn = %d want %d", got, want)
	}
}

func TestGRPCServer_Bandwidth_unavailable(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	srv := testGRPCServerUnavailable()

	req := &pb.BandwidthReq{}
	_, err := srv.Bandwidth(ctx, req)

	if got, want := errors.Cause(err), ErrUnavailable; got != want {
		t.Errorf("srv.Bandwidth(ctx, req): error = %v want %v", got, want)
	}
}