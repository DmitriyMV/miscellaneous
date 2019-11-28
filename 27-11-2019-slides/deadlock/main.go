// +build ignore,OMIT
package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
)

func main() {
	var val int64
	runtime.GOMAXPROCS(1) // HL
	go func() {
		for {
			atomic.AddInt64(&val, 1) // HL
		}
	}()
	runtime.Gosched() // HL
	fmt.Println(atomic.LoadInt64(&val))
}
