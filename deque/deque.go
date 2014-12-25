// Copyright 2014 Diego Giagio <diego@giagio.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package deque

type Deque struct {
	arr []interface{}
}

func New() *Deque {
	return &Deque{arr: nil}
}

func (d *Deque) Len() int {
	return len(d.arr)
}

func (d *Deque) PushFront(v interface{}) {
	d.arr = append([]interface{}{v}, d.arr...)
}

func (d *Deque) PushBack(v interface{}) {
	d.arr = append(d.arr, v)
}

func (d *Deque) PopFront() interface{} {
	if d.arr == nil || len(d.arr) <= 0 {
		return nil
	}
	v := d.arr[0]
	d.arr[0] = nil
	d.arr = d.arr[1:]
	return v
}

func (d *Deque) PopBack() interface{} {
	if d.arr == nil || len(d.arr) <= 0 {
		return nil
	}
	n := len(d.arr)
	v := d.arr[n-1]
	d.arr[n-1] = nil
	d.arr = d.arr[:n-1]
	return v
}

func (d *Deque) PeekFront() interface{} {
	if d.arr == nil || len(d.arr) <= 0 {
		return nil
	}
	return d.arr[0]
}

func (d *Deque) PeekBack() interface{} {
	if d.arr == nil || len(d.arr) <= 0 {
		return nil
	}
	n := len(d.arr)
	return d.arr[n-1]
}
