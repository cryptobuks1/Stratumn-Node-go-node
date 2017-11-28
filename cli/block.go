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

package cli

import (
	"context"
	"fmt"
	"io"

	"github.com/stratumn/alice/cli/script"
)

// Block is a command that evaluates a list of expressions.
var Block = BasicCmdWrapper{BasicCmd{
	Name:      "block",
	Use:       "block (command1) (command2) (etc...)",
	Short:     "Evaluate a list of expressions",
	ExecInstr: blockExec,
}}

func blockExec(
	ctx context.Context,
	cli CLI, w io.Writer,
	exec script.SExpExecutor,
	instr *script.SExp,
) error {
	for head := instr.Cdr; head != nil; head = head.Cdr {
		if head.Type != script.SExpList {
			return NewUseError("expected a list")
		}

		s, err := exec(head.List)
		if err != nil {
			return err
		}

		fmt.Fprintln(w, s)
	}

	return nil
}