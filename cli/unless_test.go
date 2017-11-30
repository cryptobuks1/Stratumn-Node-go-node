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

package cli_test

import (
	"fmt"
	"testing"

	"github.com/stratumn/alice/cli"
	"github.com/stratumn/alice/cli/script"
)

func TestUnless(t *testing.T) {
	tests := []ExecTest{{
		"unless (title test) ko ok",
		"ok\n",
		nil,
		nil,
	}, {
		"unless (test) ko ok",
		"ko\n",
		nil,
		nil,
	}, {
		"unless (title test) (title ko) (title ok)",
		"Ok\n",
		nil,
		nil,
	}, {
		"unless (test) (title ko) (title ok)",
		"Ko\n",
		nil,
		nil,
	}, {
		"unless (title test) ko",
		"",
		nil,
		nil,
	}, {
		"unless (test) ko",
		"ko\n",
		nil,
		nil,
	}, {
		"unless test ko ok",
		"ok\n",
		nil,
		nil,
	}, {
		"unless $test ko ok",
		"ko\n",
		nil,
		nil,
	}, {
		"unless test ((title o) (title k)) ((title k) (title o))",
		"O\n",
		nil,
		nil,
	}, {
		"unless $test ((title o) (title k)) ((title k) (title o))",
		"K\n",
		nil,
		nil,
	}, {
		"unless $test ((title o) k)",
		"k\n",
		nil,
		nil,
	}, {
		"unless",
		"",
		ErrUse,
		nil,
	}, {
		"unless test",
		"",
		ErrUse,
		nil,
	}, {
		"unless test ko $ok",
		"",
		script.ErrSymNotFound,
		nil,
	}, {
		"unless (test) $ko ok",
		"",
		script.ErrSymNotFound,
		nil,
	}, {
		"unless (title test)",
		"",
		ErrUse,
		nil,
	}}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d-%s", i, tt.Command), func(t *testing.T) {
			tt.Test(t, cli.Unless)
		})
	}
}
