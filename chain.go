/*
 * MIT License
 *
 * Copyright (c) 2024 RichyHBM
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package main

type Chain[T any] struct {
	Value T
	Error error
}

type chainFunction[T any] func() (T, error)
type chainFunction1[T, U any] func(args T) (U, error)
type chainFunction2[T, U, V any] func(arg1 T, arg2 U) (V, error)

func Run[T any](chainFunc chainFunction[T]) *Chain[T] {
	Value, Error := chainFunc()
	return &Chain[T]{
		Value: Value,
		Error: Error,
	}
}

func Run1[T, U any](chainFunc chainFunction1[T, U], arg T) *Chain[U] {
	Value, Error := chainFunc(arg)
	return &Chain[U]{
		Value: Value,
		Error: Error,
	}
}

func Run2[T, U, V any](chainFunc chainFunction2[T, U, V], arg1 T, arg2 U) *Chain[V] {
	Value, Error := chainFunc(arg1, arg2)
	return &Chain[V]{
		Value: Value,
		Error: Error,
	}
}

func Then[T, U any](chain *Chain[T], chainFunc chainFunction1[T, U]) *Chain[U] {
	if chain.Error == nil {
		Value, Error := chainFunc(chain.Value)
		return &Chain[U]{
			Value: Value,
			Error: Error,
		}
	} else {
		return &Chain[U]{
			Error: chain.Error,
		}
	}
}

func Result[T any](chain *Chain[T]) (T, error) {
	return chain.Value, chain.Error
}
