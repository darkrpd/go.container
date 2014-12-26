// Copyright 2014 Diego Giagio <diego@giagio.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package deque

type Deque struct {
	arr   []interface{}
	front int
	back  int
	n     int
}

func New() *Deque {
	return &Deque{arr: nil, front: 0, back: 0, n: 0}
}

func (d *Deque) PushFront(v interface{}) {
	d.growIfNeeded()

	if d.front == 0 {
		d.front = cap(d.arr) - 1
	} else {
		d.front--
	}

	d.arr[d.front] = v
	d.n++
}

func (d *Deque) PushBack(v interface{}) {
	d.growIfNeeded()

	d.arr[d.back] = v
	d.back++
	if d.back == cap(d.arr) {
		d.back = 0 // wrap-around
	}
	d.n++
}

func (d *Deque) PopFront() interface{} {
	if d.n <= 0 {
		return nil
	}

	v := d.arr[d.front]
	d.arr[d.front] = nil
	d.front++
	if d.front == cap(d.arr) {
		d.front = 0 // wrap-around
	}
	d.n--

	d.shrinkIfNeeded()
	return v
}

func (d *Deque) PopBack() interface{} {
	if d.n <= 0 {
		return nil
	}

	if d.back == 0 {
		d.back = cap(d.arr) - 1
	} else {
		d.back--
	}

	v := d.arr[d.back]
	d.arr[d.back] = nil
	d.n--

	d.shrinkIfNeeded()
	return v
}

func (d *Deque) PeekFront() interface{} {
	if d.n <= 0 {
		return nil
	}
	return d.arr[d.front]
}

func (d *Deque) PeekBack() interface{} {
	if d.n <= 0 {
		return nil
	}

	if d.back == 0 {
		return d.arr[cap(d.arr)-1]
	}
	return d.arr[d.back-1]
}

func (d *Deque) Clear() {
	d.arr = nil
	d.front = 0
	d.back = 0
	d.n = 0
}

func (d *Deque) Empty() bool {
	return d.n == 0
}

func (d *Deque) Len() int {
	return d.n
}

func (d *Deque) resize(n int) {
	newArr := make([]interface{}, n)
	for i := 0; i < d.n; i++ {
		newArr[i] = d.arr[(d.front+i)%cap(d.arr)]
	}
	d.arr = newArr
	d.front = 0
	d.back = d.n
}

func (d *Deque) growIfNeeded() {
	if d.n == cap(d.arr) {
		d.resize((d.n + 1) * 2)
	}
}

func (d *Deque) shrinkIfNeeded() {
	if d.n > 0 && d.n <= cap(d.arr)/4 {
		d.resize((cap(d.arr) - 1) / 2)
	}
}
