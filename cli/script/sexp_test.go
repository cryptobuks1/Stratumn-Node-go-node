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

package script

import (
	"strings"
	"testing"

	"github.com/pkg/errors"
)

type evalTest struct {
	input  string
	output string
	err    string
}

var evalTT = []evalTest{{
	"",
	"",
	"",
}, {
	"()",
	"",
	"",
}, {
	"echo",
	"",
	"",
}, {
	"echo hello",
	"hello",
	"",
}, {
	"echo (echo)",
	"",
	"",
}, {
	"echo echo hello",
	"echo hello",
	"",
}, {
	"(echo hello)",
	"hello",
	"",
}, {
	"(echo hello world)",
	"hello world",
	"",
}, {
	"(echo 'hello  world')",
	"hello  world",
	"",
}, {
	"(echo hello (echo world !))",
	"hello world !",
	"",
}, {
	"(echo (echo the world) (echo is beautiful) !)",
	"the world is beautiful !",
	"",
}, {
	"(+ 1 2)",
	"",
	"1:2: unknown function \"+\"",
}, {
	"(echo (+ 1 2))",
	"",
	"1:8: unknown function \"+\"",
}}

func testExecutor(sexp *SExp) (string, error) {
	if sexp == nil {
		return "", nil
	}

	if sexp.Str == "echo" {
		vals, err := sexp.Cdr.EvalEach(testExecutor)
		if err != nil {
			return "", err
		}

		return strings.Join(vals, " "), nil
	}

	return "", errors.Errorf("%d:%d: unknown function %q", sexp.Line, sexp.Offset, sexp.Str)
}

func TestSExp_eval(t *testing.T) {
	s := NewScanner()
	p := NewParser(s)

	for _, test := range evalTT {
		sexp, err := p.Parse(test.input)
		if err != nil {
			t.Errorf("%q: error: %s", test.input, err)
			continue
		}

		output, err := testExecutor(sexp)
		if err != nil {
			if test.err != "" {
				if err.Error() != test.err {
					t.Errorf("%q: error = %q want %q", test.input, err, test.err)
				}
			} else {
				t.Errorf("%q: error: %s", test.input, err)
			}
			continue
		}

		if output != test.output {
			t.Errorf("%q: output = %q want %q", test.input, output, test.output)
		}
	}
}