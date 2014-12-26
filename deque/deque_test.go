// Copyright 2014 Diego Giagio <diego@giagio.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package deque

import "testing"

func TestDeque(t *testing.T) {
	d := New()
	dbg(t, d)
	if d.Len() != 0 {
		t.Errorf("Len must be 0, it is: %d", d.Len())
	}

	d.PushBack(9)
	dbg(t, d)
	if d.Len() != 1 {
		t.Errorf("Len must be 1, it is: %d", d.Len())
	}

	if d.PeekBack() != 9 || d.PeekFront() != 9 {
		t.Errorf("PeekBack and PeekFront must be 9, it is: %d/%d", d.PeekBack(), d.PeekFront())
	}

	d.PushFront(1)
	dbg(t, d)
	if d.Len() != 2 {
		t.Errorf("Len must be 2, it is: %d", d.Len())
	}

	if d.PeekBack() != 9 || d.PeekFront() != 1 {
		t.Errorf("PeekBack must be 9 and PeekFront must be 1, it is: %d/%d", d.PeekBack(), d.PeekFront())
	}

	d.PushFront(0)
	dbg(t, d)
	if d.Len() != 3 {
		t.Errorf("Len must be 3, it is: %d", d.Len())
	}

	if d.PeekBack() != 9 || d.PeekFront() != 0 {
		t.Errorf("PeekBack must be 9 and PeekFront must be 0, it is: %d/%d", d.PeekBack(), d.PeekFront())
	}

	var v interface{}
	v = d.PopFront()
	dbg(t, d)
	if v != 0 {
		t.Errorf("PopFront must be 0, it is %d", v)
	}

	if d.Len() != 2 {
		t.Errorf("Len must be 2, it is: %d", d.Len())
	}

	if d.PeekBack() != 9 || d.PeekFront() != 1 {
		t.Errorf("PeekBack must be 9 and PeekFront must be 1, it is: %d/%d", d.PeekBack(), d.PeekFront())
	}

	v = d.PopBack()
	dbg(t, d)
	if v != 9 {
		t.Errorf("PopBack must be 9, it is %d", v)
	}

	if d.Len() != 1 {
		t.Errorf("Len must be 2, it is: %d", d.Len())
	}

	if d.PeekBack() != 1 || d.PeekFront() != 1 {
		t.Errorf("PeekBack and PeekFront must be 1, it is: %d/%d", d.PeekBack(), d.PeekFront())
	}

	v = d.PopFront()
	dbg(t, d)
	if v != 1 {
		t.Errorf("PopFront must be 1, it is %d", v)
	}

	if d.Len() != 0 {
		t.Errorf("Len must be 0, it is: %d", d.Len())
	}

	v = d.PopBack()
	dbg(t, d)
	if v != nil {
		t.Errorf("PopBack must be nil, it is: %v", v)
	}

	v = d.PopFront()
	dbg(t, d)
	if v != nil {
		t.Errorf("PopFront must be nil, it is: %v", v)
	}

	for i := 1; i <= 10; i++ {
		d.PushFront(i)
		d.PushBack(i * 2)
		dbg(t, d)
	}

	for i := 10; i >= 1; i-- {
		if n := d.PeekFront(); n != i {
			t.Errorf("PeekFront must be %d, it is: %v", i, n)
		}
		if n := d.PopFront(); n != i {
			t.Errorf("PopFront must be %d, it is: %v", i, n)
		}
		if n := d.PeekBack(); n != i*2 {
			t.Errorf("PeekBack must be %d, it is: %v", i*2, n)
		}
		if n := d.PopBack(); n != i*2 {
			t.Errorf("PopBack must be %d, it is: %v", i*2, n)
		}
		dbg(t, d)
	}
}

func dbg(t *testing.T, d *Deque) {
	t.Logf("dbg: arr=%v, n=%d, front=%d, back=%d, cap=%d\n", d.arr, d.n, d.front, d.back, cap(d.arr))
}
