package deque

import "testing"

func TestDeque(t *testing.T) {
	d := New()
	if d.Len() != 0 {
		t.Errorf("Len must be 0, it is: %d", d.Len())
	}

	d.PushBack(9)
	if d.Len() != 1 {
		t.Errorf("Len must be 1, it is: %d", d.Len())
	}

	if d.PeekBack() != 9 || d.PeekFront() != 9 {
		t.Errorf("PeekBack and PeekFront must be 9, it is: %d/%d", d.PeekBack(), d.PeekFront())
	}

	d.PushFront(1)
	if d.Len() != 2 {
		t.Errorf("Len must be 2, it is: %d", d.Len())
	}

	if d.PeekBack() != 9 || d.PeekFront() != 1 {
		t.Errorf("PeekBack must be 9 and PeekFront must be 1, it is: %d/%d", d.PeekBack(), d.PeekFront())
	}

	d.PushFront(0)
	if d.Len() != 3 {
		t.Errorf("Len must be 3, it is: %d", d.Len())
	}

	if d.PeekBack() != 9 || d.PeekFront() != 0 {
		t.Errorf("PeekBack must be 9 and PeekFront must be 0, it is: %d/%d", d.PeekBack(), d.PeekFront())
	}

	var v interface{}
	v = d.PopFront()
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
	if v != 1 {
		t.Errorf("PopFront must be 1, it is %d", v)
	}

	if d.Len() != 0 {
		t.Errorf("Len must be 0, it is: %d", d.Len())
	}

	v = d.PopBack()
	if v != nil {
		t.Errorf("PopBack must be nil, it is: %v", v)
	}

	v = d.PopFront()
	if v != nil {
		t.Errorf("PopFront must be nil, it is: %v", v)
	}
}
