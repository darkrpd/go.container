go.container
============

Elementary data structures and containers for Go language

# stack
```go
package main

import (
	"fmt"
	"github.com/dgiagio/go.container/stack"
)

func main() {
	s := stack.New()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	for s.Len() != 0 {
		v := s.Pop()
		fmt.Printf("sv=%v\n", v)
	}
}
```

# queue
```go
package main

import (
	"fmt"
	"github.com/dgiagio/go.container/queue"
)

func main() {
	q := queue.New()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	for q.Len() != 0 {
		v := q.Pop()
		fmt.Printf("qv=%v\n", v)
	}
}
```

# deque (Double-ended Queue)
```go
package main

import (
	"fmt"
	"github.com/dgiagio/go.container/deque"
)

func main() {
	d := deque.New()
	d.PushBack(2)
	d.PushBack(3)
	d.PushFront(1)
	for d.Len() != 0 {
		v := d.PopFront()
		fmt.Printf("dv=%v\n", v)
	}
}
```

# prioq (Priority Queue)
```go
package main

import (
	"fmt"
	"github.com/dgiagio/go.container/prioq"
)

func main() {
	// Comparator for int type
	comp := func(a, b interface{}) bool { return a.(int) > b.(int) }
	
	pq := prioq.New(comp)
	pq.Push(5)
	pq.Push(3)
	pq.Push(1)
	pq.Push(4)
	pq.Push(2)
	for pq.Len() != 0 {
		v := pq.Pop()
		fmt.Printf("pqv=%v\n", v)
	}
}
```
