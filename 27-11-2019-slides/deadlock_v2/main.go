// +build ignore,OMIT
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var val int64
	var mx sync.Mutex
	runtime.GOMAXPROCS(1) // HL
	go func() {
		for {
			mx.Lock() // HL
			val++
			mx.Unlock() // HL
		}
	}()
	runtime.Gosched() // HL
	mx.Lock()
	valCopy := val
	mx.Unlock()
	fmt.Println(valCopy)
}
