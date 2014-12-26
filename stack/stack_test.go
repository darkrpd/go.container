// Copyright 2014 Diego Giagio <diego@giagio.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stack

import "testing"

func TestStack(t *testing.T) {
	s := New()
	verifyLenAndCap(t, s, 0, 0)

	s.Push(1)
	verifyLenAndCap(t, s, 1, 2)

	n := s.Peek()
	verifyValue(t, n, 1)
	verifyLenAndCap(t, s, 1, 2)

	n = s.Pop()
	verifyValue(t, n, 1)
	verifyLenAndCap(t, s, 0, 2)

	n = s.Pop()
	verifyValue(t, n, nil)
	verifyLenAndCap(t, s, 0, 2)

	s.Push(9)
	verifyLenAndCap(t, s, 1, 2)
	s.Push(8)
	verifyLenAndCap(t, s, 2, 2)
	s.Push(7)
	verifyLenAndCap(t, s, 3, 6)
	s.Push(6)
	verifyLenAndCap(t, s, 4, 6)
	s.Push(5)
	verifyLenAndCap(t, s, 5, 6)
	s.Push(4)
	verifyLenAndCap(t, s, 6, 6)
	s.Push(3)
	verifyLenAndCap(t, s, 7, 14)
	s.Push(2)
	verifyLenAndCap(t, s, 8, 14)
	s.Push(1)
	verifyLenAndCap(t, s, 9, 14)

	n = s.Pop()
	verifyValue(t, n, 1)
	verifyLenAndCap(t, s, 8, 14)
	n = s.Pop()
	verifyValue(t, n, 2)
	verifyLenAndCap(t, s, 7, 14)
	n = s.Pop()
	verifyValue(t, n, 3)
	verifyLenAndCap(t, s, 6, 14)
	n = s.Pop()
	verifyValue(t, n, 4)
	verifyLenAndCap(t, s, 5, 14)
	n = s.Pop()
	verifyValue(t, n, 5)
	verifyLenAndCap(t, s, 4, 14)
	n = s.Pop()
	verifyValue(t, n, 6)
	verifyLenAndCap(t, s, 3, 6)
	n = s.Pop()
	verifyValue(t, n, 7)
	verifyLenAndCap(t, s, 2, 6)
	n = s.Pop()
	verifyValue(t, n, 8)
	verifyLenAndCap(t, s, 1, 2)
	n = s.Pop()
	verifyValue(t, n, 9)
	verifyLenAndCap(t, s, 0, 2)
	n = s.Pop()
	verifyValue(t, n, nil)
	verifyLenAndCap(t, s, 0, 2)
}

func verifyLenAndCap(t *testing.T, s *Stack, l, c int) {
	if s.Len() != l {
		t.Errorf("Len must be %d", l)
	}
	if cap(s.arr) != c {
		t.Errorf("Cap must be %d", c)
	}
}

func verifyValue(t *testing.T, n, expected interface{}) {
	if n != expected {
		t.Errorf("Value incorrect: %d, expected: %d", n, expected)
	}
}
