// Copyright 2014 Diego Giagio <diego@giagio.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package queue

type Queue struct {
	arr   []interface{}
	front int
	back  int
	n     int
}

func New() *Queue {
	return &Queue{arr: nil, front: 0, back: 0, n: 0}
}

func (q *Queue) Push(v interface{}) {
	// Grow capacity if needed
	if q.n == cap(q.arr) {
		q.resize((q.n + 1) * 2)
	}

	q.arr[q.back] = v
	q.back++
	if q.back == cap(q.arr) {
		q.back = 0 // wrap-around
	}
	q.n++
}

func (q *Queue) Pop() interface{} {
	if q.n <= 0 {
		return nil
	}

	v := q.arr[q.front]
	q.arr[q.front] = nil
	q.front++
	if q.front == cap(q.arr) {
		q.front = 0 // wrap-around
	}
	q.n--

	// Shrink if capacity is too big
	if q.n > 0 && q.n <= cap(q.arr)/4 {
		q.resize((cap(q.arr) - 1) / 2)
	}

	return v
}

func (q *Queue) Peek() interface{} {
	if q.n <= 0 {
		return nil
	}
	return q.arr[q.front]
}

func (q *Queue) Len() int {
	return q.n
}

func (q *Queue) resize(n int) {
	newArr := make([]interface{}, n)
	for i := 0; i < q.n; i++ {
		newArr[i] = q.arr[(q.front+i)%cap(q.arr)]
	}
	q.arr = newArr
	q.front = 0
	q.back = q.n
}
