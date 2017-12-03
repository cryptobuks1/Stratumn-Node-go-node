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

	"github.com/stratumn/alice/cli/script"
)

// Eval is a command that evaluates the contents of a string.
var Eval = BasicCmdWrapper{BasicCmd{
	Name:     "eval",
	Use:      "eval <Expression>",
	Short:    "Evaluate an expression",
	ExecSExp: evalExec,
}}

func evalExec(
	ctx context.Context,
	cli CLI,
	closure *script.Closure,
	call script.CallHandler,
	args script.SCell,
	meta script.Meta,
) (script.SExp, error) {
	// Eval is a simple command, all it does is evaluate the car twice.

	if args == nil || args.Cdr() != nil {
		return nil, NewUseError("expected a single expression")
	}

	v, err := script.Eval(closure.Resolve, call, args.Car())
	if err != nil {
		return nil, err
	}

	return script.Eval(closure.Resolve, call, v)
}
