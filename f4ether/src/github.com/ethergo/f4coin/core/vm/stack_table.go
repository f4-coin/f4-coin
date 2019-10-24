// Copyright 2017 The f4coin Authors
// This file is part of the f4coin library.
//
// The go-f4coin library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-f4coin library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with The go-f4coin library. If not, see <http://www.gnu.org/licenses/>.

package vm

import (
	"fmt"

	"github.com/ethergo/f4coin/params"
)

func makeStackFunc(pop, push int) stackValidationFunc {
	return func(stack *Stack) error {
		if err := stack.require(pop); err != nil {
			return err
		}

		if stack.len()+push-pop > int(params.StackLimit) {
			return fmt.Errorf("stack limit reached %d (%d)", stack.len(), params.StackLimit)
		}
		return nil
	}
}

func makeDupStackFunc(n int) stackValidationFunc {
	return makeStackFunc(n, n+1)
}

func makeSwapStackFunc(n int) stackValidationFunc {
	return makeStackFunc(n, n)
}
