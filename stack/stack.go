// Copyright 2014 Diego Giagio <diego@giagio.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stack

type Stack struct {
	arr []interface{}
	n   int
}

func New() *Stack {
	return &Stack{arr: nil, n: 0}
}

func (s *Stack) Push(v interface{}) {
	s.growIfNeeded()
	s.arr[s.n] = v
	s.n++
}

func (s *Stack) Pop() interface{} {
	if s.n <= 0 {
		return nil
	}

	s.n--
	v := s.arr[s.n]
	s.arr[s.n] = nil

	s.shrinkIfNeeded()
	return v
}

func (s *Stack) Peek() interface{} {
	if s.n <= 0 {
		return nil
	}
	return s.arr[s.n-1]
}

func (s *Stack) Clear() {
	s.arr = nil
	s.n = 0
}

func (s *Stack) Empty() bool {
	return s.n == 0
}

func (s *Stack) Len() int {
	return s.n
}

func (s *Stack) resize(n int) {
	newArr := make([]interface{}, n)
	copy(newArr, s.arr)
	s.arr = newArr
}

func (s *Stack) growIfNeeded() {
	if s.n == cap(s.arr) {
		s.resize((s.n + 1) * 2)
	}
}

func (s *Stack) shrinkIfNeeded() {
	if s.n > 0 && s.n <= cap(s.arr)/4 {
		s.resize((cap(s.arr) - 1) / 2)
	}
}
