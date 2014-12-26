// Copyright 2014 Diego Giagio <diego@giagio.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package queue

import "testing"

func TestQueue(t *testing.T) {
	q := New()
	dbg(t, q)
	verifyLenAndCap(t, q, 0, 0)

	q.Push(1)
	dbg(t, q)
	verifyLenAndCap(t, q, 1, 2)

	n := q.Peek()
	verifyValue(t, n, 1)
	verifyLenAndCap(t, q, 1, 2)

	n = q.Pop()
	dbg(t, q)
	verifyValue(t, n, 1)
	verifyLenAndCap(t, q, 0, 2)

	n = q.Pop()
	dbg(t, q)
	verifyValue(t, n, nil)
	verifyLenAndCap(t, q, 0, 2)

	q.Push(9)
	dbg(t, q)
	verifyLenAndCap(t, q, 1, 2)
	q.Push(8)
	dbg(t, q)
	verifyLenAndCap(t, q, 2, 2)
	q.Push(7)
	dbg(t, q)
	verifyLenAndCap(t, q, 3, 6)
	q.Push(6)
	dbg(t, q)
	verifyLenAndCap(t, q, 4, 6)
	q.Push(5)
	dbg(t, q)
	verifyLenAndCap(t, q, 5, 6)
	q.Push(4)
	dbg(t, q)
	verifyLenAndCap(t, q, 6, 6)
	q.Push(3)
	dbg(t, q)
	verifyLenAndCap(t, q, 7, 14)
	q.Push(2)
	dbg(t, q)
	verifyLenAndCap(t, q, 8, 14)
	q.Push(1)
	dbg(t, q)
	verifyLenAndCap(t, q, 9, 14)

	n = q.Pop()
	dbg(t, q)
	verifyValue(t, n, 9)
	verifyLenAndCap(t, q, 8, 14)
	n = q.Pop()
	dbg(t, q)
	verifyValue(t, n, 8)
	verifyLenAndCap(t, q, 7, 14)
	n = q.Pop()
	dbg(t, q)
	verifyValue(t, n, 7)
	verifyLenAndCap(t, q, 6, 14)
	n = q.Pop()
	dbg(t, q)
	verifyValue(t, n, 6)
	verifyLenAndCap(t, q, 5, 14)
	n = q.Pop()
	dbg(t, q)
	verifyValue(t, n, 5)
	verifyLenAndCap(t, q, 4, 14)
	n = q.Pop()
	dbg(t, q)
	verifyValue(t, n, 4)
	verifyLenAndCap(t, q, 3, 6)
	n = q.Pop()
	dbg(t, q)
	verifyValue(t, n, 3)
	verifyLenAndCap(t, q, 2, 6)
	n = q.Pop()
	dbg(t, q)
	verifyValue(t, n, 2)
	verifyLenAndCap(t, q, 1, 2)
	n = q.Pop()
	dbg(t, q)
	verifyValue(t, n, 1)
	verifyLenAndCap(t, q, 0, 2)
	n = q.Pop()
	dbg(t, q)
	verifyValue(t, n, nil)
	verifyLenAndCap(t, q, 0, 2)

	q.Push(1)
	dbg(t, q)
	verifyLenAndCap(t, q, 1, 2)
	q.Push(2)
	dbg(t, q)
	verifyLenAndCap(t, q, 2, 2)
	n = q.Pop()
	dbg(t, q)
	verifyValue(t, n, 1)
	verifyLenAndCap(t, q, 1, 2)
	n = q.Pop()
	dbg(t, q)
	verifyValue(t, n, 2)
	verifyLenAndCap(t, q, 0, 2)

	q.Push(1)
	dbg(t, q)
	verifyLenAndCap(t, q, 1, 2)
	q.Push(2)
	dbg(t, q)
	verifyLenAndCap(t, q, 2, 2)
	q.Push(3)
	dbg(t, q)
	verifyLenAndCap(t, q, 3, 6)
	n = q.Pop()
	dbg(t, q)
	verifyValue(t, n, 1)
	verifyLenAndCap(t, q, 2, 6)
	n = q.Pop()
	dbg(t, q)
	verifyValue(t, n, 2)
	verifyLenAndCap(t, q, 1, 2)
	n = q.Pop()
	dbg(t, q)
	verifyValue(t, n, 3)
	verifyLenAndCap(t, q, 0, 2)
}

func dbg(t *testing.T, q *Queue) {
	t.Logf("dbg: arr=%v, n=%d, front=%d, back=%d, cap=%d\n", q.arr, q.n, q.front, q.back, cap(q.arr))
}

func verifyLenAndCap(t *testing.T, q *Queue, l, c int) {
	if q.Len() != l {
		t.Errorf("Len must be %d", l)
	}
	if len(q.arr) != c {
		t.Errorf("Cap must be %d", c)
	}
}

func verifyValue(t *testing.T, n, expected interface{}) {
	if n != expected {
		t.Errorf("Value incorrect: %d, expected: %d", n, expected)
	}
}
