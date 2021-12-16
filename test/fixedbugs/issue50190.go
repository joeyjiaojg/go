// run

// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

type int float32

type Int = int

type A = struct{ int }
type B = struct{ Int }

func main() {
	var x, y interface{} = A{}, B{}
	if x == y {
		panic("FAIL")
	}
}