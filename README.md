# Data structures with Go

This is a collection of some data structures with Go. They're built with TDD.

## How to use this module
- Init a Go Module (e.g. `go mod init example.com`)
- `go get github.com/unpolinomio/data-structures-with-go`
- Have fun!

## Example
```go
package main

import (
	"fmt"

	datas "github.com/unpolinomio/data-structures-with-go"
)

func main() {
	list := datas.SynglyLinkedList{}
	list.PushFront(10)
	list.PushFront(20)
	list.PushFront(30)

	fmt.Println(list.Get(1))
}
```

## Data structures included
- Singly-linked list
- ...
