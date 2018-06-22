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

package dummyauditstore_test

import (
	"testing"

	"github.com/stratumn/alice/app/indigo/protocol/store/audit"
	"github.com/stratumn/alice/app/indigo/protocol/store/audit/dummyauditstore"
	"github.com/stratumn/alice/app/indigo/protocol/store/audit/storetestcases"
)

func TestDummyAuditStore(t *testing.T) {
	storetestcases.Factory{
		New: func() (audit.Store, error) {
			return dummyauditstore.New(), nil
		},
		Free: func(audit.Store) {},
	}.RunTests(t)
}