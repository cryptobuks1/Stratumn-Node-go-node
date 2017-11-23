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
	"strings"
	"testing"
	"time"

	"github.com/spf13/pflag"
)

type basicContentMock string

func (s basicContentMock) TextBeforeCursor() string {
	return string(s)
}

func (s basicContentMock) GetWordBeforeCursor() string {
	parts := strings.Split(string(s), " ")
	if len(parts) < 1 {
		return ""
	}
	return parts[len(parts)-1]
}

func TestBasicCmdWrapper_strings(t *testing.T) {
	cmd := BasicCmdWrapper{BasicCmd{
		Name:  "cmd",
		Short: "A test command",
	}}

	if got, want := cmd.Name(), "cmd"; got != want {
		t.Errorf("cmd.Name() = %v want %v", got, want)
	}

	if got, want := cmd.Short(), "A test command"; got != want {
		t.Errorf("cmd.Short() = %v want %v", got, want)
	}

	if got, want := cmd.Long(), "A test command\n\nUsage:\n  cmd\n\nFlags:\n  -h, --help   Invoke help on command\n"; got != want {
		t.Errorf("cmd.Long() = \n%v want\n%v", got, want)
	}
}

func TestBasicCmdWrapper_Suggest(t *testing.T) {
	cmd := BasicCmdWrapper{BasicCmd{
		Name:  "cmd",
		Short: "A test command",
		Flags: func() *pflag.FlagSet {
			flags := pflag.NewFlagSet("cmd", pflag.ContinueOnError)
			flags.String("flag", "", "")
			return flags
		},
		Exec: func(context.Context, *CLI, []string, *pflag.FlagSet) error {
			return nil
		},
	}}

	tt := []struct {
		name   string
		text   string
		expect []string
	}{{
		"empty",
		"",
		[]string{"cmd"},
	}, {
		"partial",
		"cm",
		[]string{"cmd"},
	}, {
		"mismatch",
		"cdm",
		[]string{},
	}, {
		"empty flag",
		"cmd -",
		[]string{"--flag", "--help"},
	}, {
		"partial flag",
		"cmd --f",
		[]string{"--flag"},
	}, {
		"flag mismatch",
		"cmd --fgas",
		[]string{},
	}}

	for _, test := range tt {
		content := basicContentMock(test.text)
		suggs := cmd.Suggest(content)
		completions := make([]string, len(suggs))
		for i, s := range suggs {
			completions[i] = s.Text
		}

		got := fmt.Sprintf("%s", completions)
		want := fmt.Sprintf("%s", test.expect)

		if got != want {
			t.Errorf("%s: completions = %v want %v", test.name, got, want)
		}
	}
}

func TestBasicCmdWrapper_Exec(t *testing.T) {
	execCh := make(chan struct{})

	cmd := BasicCmdWrapper{BasicCmd{
		Name:  "cmd",
		Short: "A test command",
		Exec: func(context.Context, *CLI, []string, *pflag.FlagSet) error {
			close(execCh)
			return nil
		},
	}}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if err := cmd.Exec(ctx, nil, ""); err != nil {
		t.Errorf(`cmd.Exec(ctx, nil, ""): error: %s`, err)
	}

	select {
	case <-execCh:
	case <-ctx.Done():
		t.Errorf("command was not executed")
	}
}