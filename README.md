# Worker
## Installation
To install this package, you need to setup your Go workspace.  The simplest way to install the library is to run:
```
$ go get github.com/namcuongq/worker
```
## Example
```go
package main

import (
	"fmt"
	"github.com/namcuongq/worker"
)

func Tak(data interface{}){
  fmt.Println(data)
}

func main() {
	pool := worker.New(2)
	requests := []string{"alpha", "beta", "gamma", "delta", "epsilon"}

	for _, r := range requests {
		pool.Add(r)
	}

	pool.WaitAndClose()
}
