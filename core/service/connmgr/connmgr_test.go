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

package connmgr

import (
	"context"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/stratumn/alice/core/manager/testservice"

	ifconnmgr "gx/ipfs/QmYkCrTwivapqdB3JbwvwvxymseahVkcm46ThRMAA24zCr/go-libp2p-interface-connmgr"
)

func testService(ctx context.Context, t *testing.T) *Service {
	serv := &Service{}
	config := serv.Config().(Config)

	if err := serv.SetConfig(config); err != nil {
		t.Fatalf("serv.SetConfig(config): error: %s", err)
	}

	return serv
}

func TestService_strings(t *testing.T) {
	testservice.CheckStrings(t, &Service{})
}

func TestService_Expose(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	serv := testService(ctx, t)
	exposed := testservice.Expose(ctx, t, serv, time.Second)

	_, ok := exposed.(ifconnmgr.ConnManager)
	if got, want := ok, true; got != want {
		t.Errorf("ok = %v want %v", got, want)
	}
}

func TestService_Run(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	serv := testService(ctx, t)
	testservice.TestRun(ctx, t, serv, time.Second)
}

func TestService_SetConfig(t *testing.T) {
	errAny := errors.New("any error")

	tests := []struct {
		name string
		set  func(*Config)
		err  error
	}{{
		"invalid grace period",
		func(c *Config) { c.GracePeriod = "1" },
		errAny,
	}}

	for _, tt := range tests {
		serv := Service{}
		config := serv.Config().(Config)
		tt.set(&config)

		err := errors.Cause(serv.SetConfig(config))
		switch {
		case err != nil && tt.err == errAny:
		case err != tt.err:
			t.Errorf("%s: err = %v want %v", tt.name, err, tt.err)
		}
	}
}
