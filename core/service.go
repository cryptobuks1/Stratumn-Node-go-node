// Copyright © 2017 Stratumn SAS
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

package core

import (
	"github.com/stratumn/alice/core/manager"
	"github.com/stratumn/alice/core/service/bootstrap"
	"github.com/stratumn/alice/core/service/connmgr"
	"github.com/stratumn/alice/core/service/grpcapi"
	"github.com/stratumn/alice/core/service/host"
	"github.com/stratumn/alice/core/service/identify"
	"github.com/stratumn/alice/core/service/kaddht"
	"github.com/stratumn/alice/core/service/metrics"
	"github.com/stratumn/alice/core/service/mssmux"
	"github.com/stratumn/alice/core/service/natmgr"
	"github.com/stratumn/alice/core/service/ping"
	"github.com/stratumn/alice/core/service/pruner"
	"github.com/stratumn/alice/core/service/relay"
	"github.com/stratumn/alice/core/service/signal"
	"github.com/stratumn/alice/core/service/swarm"
	"github.com/stratumn/alice/core/service/yamux"
)

// services contains all the services.
var services = []manager.Service{
	&grpcapi.Service{},
	&pruner.Service{},
	&signal.Service{},
	&yamux.Service{},
	&mssmux.Service{},
	&swarm.Service{},
	&connmgr.Service{},
	&host.Service{},
	&natmgr.Service{},
	&metrics.Service{},
	&relay.Service{},
	&identify.Service{},
	&bootstrap.Service{},
	&kaddht.Service{},
	&ping.Service{},
}
