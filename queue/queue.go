// Copyright 2014 Diego Giagio <diego@giagio.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package queue

import "github.com/dgiagio/go.container/deque"

type Queue struct {
	deque.Deque
}

func New() *Queue {
	return &Queue{deque.Deque{}}
}

func (q *Queue) Push(v interface{}) {
	q.Deque.PushBack(v)
}

func (q *Queue) Pop() interface{} {
	return q.Deque.PopFront()
}

func (q *Queue) Peek() interface{} {
	return q.Deque.PeekFront()
}

func (q *Queue) Len() int {
	return q.Deque.Len()
}
