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

package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
)

// installCmd represents the install command.
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Register the daemon service",
	Run: func(cmd *cobra.Command, args []string) {
		cfgPath, err := filepath.Abs(coreCfgFilename)
		fail(err)

		status, err := newDaemon().Install("up", "--core-config", cfgPath)
		fail(err)

		fmt.Println(status)
	},
}

func init() {
	daemonCmd.AddCommand(installCmd)
}