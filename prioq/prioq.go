// Copyright 2014 Diego Giagio <diego@giagio.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package prioq

import "container/heap"

type Comparator func(a, b interface{}) bool

type prioHeap struct {
	comp Comparator
	arr  []interface{}
}

func (h *prioHeap) Len() int {
	return len(h.arr)
}

func (h *prioHeap) Less(i, j int) bool {
	return h.comp(h.arr[i], h.arr[j])
}

func (h *prioHeap) Swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

func (h *prioHeap) Push(x interface{}) {
	h.arr = append(h.arr, x)
}

func (h *prioHeap) Pop() interface{} {
	n := len(h.arr)
	v := h.arr[n-1]
	h.arr = h.arr[:n-1]
	return v
}

type PriorityQueue struct {
	h prioHeap
}

func New(comp Comparator) *PriorityQueue {
	pq := &PriorityQueue{h: prioHeap{comp: comp, arr: nil}}
	heap.Init(&pq.h)
	return pq
}

func (pq *PriorityQueue) Push(v interface{}) {
	heap.Push(&pq.h, v)
}

func (pq *PriorityQueue) Pop() interface{} {
	return heap.Pop(&pq.h)
}

func (pq *PriorityQueue) Len() int {
	return len(pq.h.arr)
}
