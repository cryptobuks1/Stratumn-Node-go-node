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
	"context"
	"fmt"
	"testing"
)

func BenchmarkInterpreter_reverse_rec_tail_opt(b *testing.B) {
	for i := 10; i <= 10000; i *= 10 {
		b.Run(fmt.Sprintf("depth-%d", i), func(b *testing.B) {
			benchmarkReverseRec(b, i, true)
		})
	}
}

func BenchmarkInterpreter_reverse_rec_no_tail_opt(b *testing.B) {
	for i := 10; i <= 10000; i *= 10 {
		b.Run(fmt.Sprintf("depth-%d", i), func(b *testing.B) {
			benchmarkReverseRec(b, i, false)
		})
	}
}

func benchmarkReverseRec(b *testing.B, depth int, tailOptimize bool) {
	src := `
		; Reverses a list recusively.
		let reverse (lambda (l) (
			; Define a nested recursive function with an accumulator.
			(let reverse-rec (lambda (l tail) (
				(unless l 
					tail
					(reverse-rec (cdr l) (cons (car l) tail))))))
			; Start the recursion
			(reverse-rec l ())))
		let list (quote(`
	for i := 0; i < depth; i++ {
		src += fmt.Sprint(i) + " "
	}
	src += "))"

	itr := NewInterpreter(
		InterpreterOptBuiltinLibs,
		InterpreterOptTailOptimizations(tailOptimize),
		InterpreterOptValueHandler(func(SExp) {}),
	)

	err := itr.EvalInput(context.Background(), src)
	if err != nil {
		b.Fatalf(err.Error())
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := itr.EvalInput(context.Background(), "reverse list")
		if err != nil {
			b.Fatalf(err.Error())
		}
	}
}