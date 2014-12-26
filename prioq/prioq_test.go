// Copyright 2014 Diego Giagio <diego@giagio.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package prioq

import "testing"

type entry struct {
	v int
}

func prioqMin(a, b interface{}) bool {
	return a.(entry).v < b.(entry).v
}

func prioqMax(a, b interface{}) bool {
	return a.(entry).v > b.(entry).v
}

func TestPriorityQueue(t *testing.T) {
	pq := New(prioqMin)
	pq.Push(entry{v: 1})
	pq.Push(entry{v: 2})
	pq.Push(entry{v: 4})
	pq.Push(entry{v: 9})
	pq.Push(entry{v: 15})
	pq.Push(entry{v: 13})
	pq.Push(entry{v: 12})
	pq.Push(entry{v: 6})
	pq.Push(entry{v: 0})
	pq.Push(entry{v: 99})
	pq.Push(entry{v: 32})
	pq.Push(entry{v: 59})
	pq.Push(entry{v: 20})
	for pq.Len() != 0 {
		t.Logf("Len=%d\n", pq.Len())
		v := pq.Pop()
		t.Logf("v=%v\n", v)
	}

	t.Log("--\n")

	pq = New(prioqMax)
	pq.Push(entry{v: 1})
	pq.Push(entry{v: 2})
	pq.Push(entry{v: 4})
	pq.Push(entry{v: 9})
	pq.Push(entry{v: 15})
	pq.Push(entry{v: 13})
	pq.Push(entry{v: 12})
	pq.Push(entry{v: 6})
	pq.Push(entry{v: 0})
	pq.Push(entry{v: 99})
	pq.Push(entry{v: 32})
	pq.Push(entry{v: 59})
	pq.Push(entry{v: 20})
	for pq.Len() != 0 {
		t.Logf("Len=%d\n", pq.Len())
		v := pq.Pop()
		t.Logf("v=%v\n", v)
	}
}
