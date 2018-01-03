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

package core

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/pelletier/go-toml"
	logger "github.com/stratumn/alice/core/log"
	"github.com/stratumn/alice/core/service/kaddht"
	"github.com/stratumn/alice/core/service/swarm"
)

func TestMigrations(t *testing.T) {
	dir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatalf(`ioutil.TempDir("", ""): error: %s`, err)
	}

	filename := filepath.Join(dir, "cfg.toml")

	// Save original configuration.
	if err := ioutil.WriteFile(filename, []byte(confZero), 0600); err != nil {
		t.Fatalf("ioutil.WriteFile(filename, []byte(confZero), 0600): error: %s", err)
	}

	set := NewConfigurableSet(BuiltinServices())

	// Migrate and load.
	if err := LoadConfig(set, filename); err != nil {
		t.Fatalf("LoadConfig(set, filename): error: %s", err)
	}

	migratedConf := set.Configs()

	// Create default config.
	defConf := NewConfigurableSet(BuiltinServices()).Configs()

	// Make sure both configs use the same private key and point to the
	// same files.
	swarmCfg := defConf["swarm"].(swarm.Config)
	swarmCfg.PeerID = confZeroPID
	swarmCfg.PrivateKey = confZeroPK
	defConf["swarm"] = swarmCfg
	dhtCfg := defConf["kaddht"].(kaddht.Config)
	dhtCfg.LevelDBPath = "data/kaddht"
	defConf["kaddht"] = dhtCfg
	logCfg := defConf["log"].(logger.Config)
	logCfg.Writers[0].Filename = "log.jsonld"
	defConf["log"] = logCfg

	// If migrations are consistent, both configs should be the same.
	for k, gotVal := range migratedConf {
		gotBuf := bytes.NewBuffer(nil)
		enc := toml.NewEncoder(gotBuf)
		enc.QuoteMapKeys(true)
		if err := enc.Encode(gotVal); err != nil {
			t.Errorf("%s: enc.Encode(gotVal): error: %s", k, err)
			continue
		}

		wantVal := defConf[k]
		wantBuf := bytes.NewBuffer(nil)
		enc = toml.NewEncoder(wantBuf)
		enc.QuoteMapKeys(true)
		if err := enc.Encode(wantVal); err != nil {
			t.Errorf("%s: enc.Encode(wantVal): error: %s", k, err)
			continue
		}

		if got, want := gotBuf.String(), wantBuf.String(); got != want {
			t.Errorf("%s: got:\n%s\nwant\n%s\n", k, got, want)
		}
	}
}

const confZeroPID = "QmRK6DGwgkzpHRG5rCgWUxTSKvQtmhZeMa5jdeHYTtxnNq"

const confZeroPK = "CAASpwkwggSjAgEAAoIBAQCss3uW+aVGR5JLd3uw2Wubwm6nGV3/R2DFG8meu4NKB3PvmmmsngjCok++h0J6/jkqKE6o4uDubbP891GqjSOB9U/NrL8WYJV0xOKzuLgaJ3+alv3Fltg/hFH38GsMycI/3elBQpGWz56xYk/Jd+kUFxUdC9IaqC6oVdDDQHM5ItmCvDygTWTazPh1OFeqFfYjjH0GHFFHOhzRQvAA/CWFll6it7ZM+Oca/MLWwGxiiMvNKkwhZnL7oYJh4NIqAnJBSwHSAusKIBNLpo28gDbvMKuaCDLp97qqLB0tzD0wFVBsBOYbdRwyuKnR30TTIYmcunIIz6xjt2W76IzHiXIHAgMBAAECggEAHoihKkE7ImDXEbavTMY0C1bl/28xukehaVgPXpwiiz0kF1HCqz1JqTfPR41ciKhX7QcxWRS71gvZTblgW/oWNAzuLhwDsO4knn+M4V+gVSd0nR1jAsM3uosnfuGn25v0Vxxh+CLP4M0WbqBGIQWtVNr75aXIYOQpU6PQhCTp+kjPs0ibIqpc+KEHSwQvHNbE7N0oeoikQGhkX8lKGvPw0Q7jvuEr/C/S8vzciv6GyFkw2LncpvXOIBGNlXT5ZrPzLP1QrGMp3DxxTAWK0IzIwkjYahsURRQFbWxJkqcbKw1sLT/hwMAvMMuT4O5z79RwoJR/ril3r6vptXWZV5lH2QKBgQDXFGQU8LEVsmWhOpYEaqQCGMjjSWWzSGm7Yd60eJydEUpvf4l3VlZs+8JjFKJGyMyp/eRo+7J3A0JcsOgrJLaR/faaTSIIh3FDcOVez6Yl4bqQUlTVwBH4AhkwTkqK1QQ2hNThbvXbv093AhxsY3f3SGg8fkuN2ziFeKzq/Wh3FQKBgQDNjwOjI8S0Xo4u6DEdNMEezS/f00S058zZtZziHTIMEQ7fFJQZ3DFVAjkgBWRBVLb1NEq9U6XdTpD/lFNUBY3DH2qjWv4LBYPk5kqFgiuKnFmoD/gnzchGs2+4EgUdODZB/QPPBCqXA7XN/JT7VoiTWGIzAf7lv84YgUmT2BQLqwKBgCdSlRG3B8ldunMF0ROxo5a2jVPwwWVL4fjeZec8/fVBighkmu90m4yFYv7WcOzcHX8e6jm/etuDfwiPV4M7zR1X/1QqsgQ5Lx4Tb/wrnsbiREfKpbQGz8I2MADC76H+XCzTkFA/BzhL++1YN3YhoXdWh6g3tvySjfzpGURFXGoZAoGAGJS6jZ6wXhVUkV1oyiJN2b4VtIFSHQP/JiWmng95taGwkpKmZzVCnPTIGgErDPjxa/8V1PAUzJMhmb6F/G0xl5zBJsmxyWWecRfs32xCgq/RtNw8A56DDZlVicB15hmbu2ZjNzU7VpW1/uzub+PYLy6Jh6n8bkLyhVGol8pmE0MCgYEArcC+amAbHBU5DExuoXCP6pF/5Q6o13tFuG3PwWoJMCFgIGDcckBl4Hs102tkdhkH+KVmwhUFPIUp8Dg9GiEkhZi8PkkhVhAJwJEIZaNSNyOja1b5K8i7wm9xJWVumLD7wPz2I/dHv/jfHFZ1KsCb9yZg4aqGWVIRt8+vgIb97mI="

// Original configuration before migrations.
const confZero = `
# Alice configuration file. Keep private!!!

# Settings for the bootstrap module.
[bootstrap]

  # A list of known peer addresses.
  addresses = ["/dnsaddr/impulse.io/ipfs/Qmc1QbSba7RtPgxEw4NqXNeDpB5CpCTwv9dvdZRdTkche1","/dnsaddr/impulse.io/ipfs/QmQVdocY8ZbYxrKRSrff2Vxmm27Mhu6DgXyWXQwmuz1b6P","/dnsaddr/impulse.io/ipfs/QmQJib6mnEMgdCe3bGH1YP7JswHbQQejyNucvW9BjFqmWr","/dnsaddr/impulse.io/ipfs/Qmc1rLFp5stHrjtq4duFg6KakBcDCpB3bTjjMZVSAdnHLj"]

  # The connection timeout. It should be less than the interval.
  connection_timeout = "10s"

  # The name of the host service.
  host = "host"

  # Interval between bootstrap jobs.
  interval = "30s"

  # The number of peers under which to bootstrap connections.
  min_peer_threshold = 4

  # Services that should be started in addition to the host before bootstrapping.
  needs = ["network","p2p"]

# Settings for the clock module.
[clock]

  # The name of the host service.
  host = "host"

  # How long to wait before closing the stream when writing the time to a peer.
  write_timeout = "10s"

# Settings for the connmgr module.
[connmgr]

  # How long to keep a connection before it can be closed.
  grace_period = "20s"

  # Maximum number of connections to keep open (0 = disabled).
  high_water = 900

  # Minimum number of connections to keep open (0 = disabled).
  low_water = 600

# Settings for the core module.
[core]

  # Service to launch when starting the node.
  boot_service = "boot"

  # Whether to show the boot screen when starting the node.
  enable_boot_screen = true

  # The version of Alice that generated this file.
  generated_by_version = "v0.0.1"

  [[core.service_groups]]

    # Description of the service group.
    description = "Starts boot services."

    # Unique identifier of the service group.
    id = "boot"

    # Name of the service group.
    name = "Boot Services"

    # Services started by the group.
    services = ["system","bootstrap","api"]

  [[core.service_groups]]

    # Description of the service group.
    description = "Starts system services."

    # Unique identifier of the service group.
    id = "system"

    # Name of the service group.
    name = "System Services"

    # Services started by the group.
    services = ["signal","pruner"]

  [[core.service_groups]]

    # Description of the service group.
    description = "Starts P2P services."

    # Unique identifier of the service group.
    id = "p2p"

    # Name of the service group.
    name = "P2P Services"

    # Services started by the group.
    services = ["identify","relay","kaddht","ping","clock"]

  [[core.service_groups]]

    # Description of the service group.
    description = "Starts network services."

    # Unique identifier of the service group.
    id = "network"

    # Name of the service group.
    name = "Network Services"

    # Services started by the group.
    services = ["host","natmgr"]

  [[core.service_groups]]

    # Description of the service group.
    description = "Starts API services."

    # Unique identifier of the service group.
    id = "api"

    # Name of the service group.
    name = "API Services"

    # Services started by the group.
    services = ["grpcapi"]

# Settings for the grpcapi module.
[grpcapi]

  # Address to bind to.
  address = "/ip4/127.0.0.1/tcp/8904"

  # Whether to log requests.
  enable_request_logger = true

  # The name of the manager service.
  manager = "manager"

  # Path to a TLS certificate.
  tls_certificate_file = ""

  # Path to a TLS key.
  tls_key_file = ""

# Settings for the host module.
[host]

  # CIDR netmasks to filter announced addresses.
  addresses_netmasks = []

  # The name of the connection manager service.
  connection_manager = "connmgr"

  # The name of the metrics service (blank = disabled).
  metrics = "metrics"

  # The negotiation timeout.
  negotiation_timeout = "1m"

  # The name of the network or swarm service.
  network = "swarm"

# Settings for the identify module.
[identify]

  # The name of the host service.
  host = "host"

# Settings for the kaddht module.
[kaddht]

  # The name of the bootstrap service.
  bootstrap = "bootstrap"

  # How long to wait between bootstrap jobs.
  bootstrap_interval = "1m"

  # The number of queries during a bootstrap job.
  bootstrap_queries = 3

  # The timeout for a bootstrap job.
  bootstrap_timeout = "10s"

  # Whether to run only as a client and not store DHT values.
  enable_client_mode = false

  # The name of the host service.
  host = "host"

  # The path to the LevelDB database directory.
  level_db_path = "data/kaddht"

# Settings for the log module.
[log]

  [[log.writers]]

    # Whether to compress the file.
    compress = false

    # The file for a file logger.
    filename = "log.jsonld"

    # The formatter for the writer (json, text, color, journald).
    formatter = "json"

    # The log level for the writer (info, error, all).
    level = "all"

    # The maximum age of the file in days before a rotation.
    maximum_age = 7

    # The maximum number of backups.
    maximum_backups = 4

    # The maximum size of the file in megabytes before a rotation.
    maximum_size = 128

    # The type of writer (file, stdout, stderr).
    type = "file"

    # Whether to use local time instead of UTC for backups.
    use_local_time = false

# Settings for the metrics module.
[metrics]

  # Interval between updates of periodic stats.
  interval = "10s"

  # Address of the endpoint to expose Prometheus metrics (blank = disabled).
  prometheus_endpoint = "/ip4/127.0.0.1/tcp/8905"

# Settings for the mssmux module.
[mssmux]

  # A map of protocols to stream muxers (protocol = service).
  [mssmux.routes]
    "/yamux/v1.0.0" = "yamux"

# Settings for the natmgr module.
[natmgr]

  # The name of the host service.
  host = "host"

# Settings for the ping module.
[ping]

  # The name of the host service.
  host = "host"

# Settings for the pruner module.
[pruner]

  # Interval between prune jobs.
  interval = "1m"

  # The name of the manager service.
  manager = "manager"

# Settings for the relay module.
[relay]

  # Whether to act as an intermediary node in relay circuits.
  enable_hop = false

  # The name of the host service.
  host = "host"

# Settings for the signal module.
[signal]

  # Allow forced shutdown by sending second signal.
  allow_forced_shutdown = true

  # The name of the manager service.
  manager = "manager"

# Settings for the swarm module.
[swarm]

  # List of addresses to bind to.
  addresses = ["/ip4/0.0.0.0/tcp/8903","/ip6/::/tcp/8903"]

  # The name of the metrics service (blank = disabled).
  metrics = "metrics"

  # The peer ID of the host.
  peer_id = "` + confZeroPID + `"

  # The private key of the host.
    private_key = "` + confZeroPK + `"

  # The name of the stream muxer service.
  stream_muxer = "mssmux"

# Settings for the yamux module.
[yamux]

  # The size of the accept backlog.
  accept_backlog = 512

  # The connection write timeout.
  connection_write_timeout = "10s"

  # The keep alive interval.
  keep_alive_interval = "30s"

  # The maximum stream window size.
  max_stream_window_size = "512KB"
`