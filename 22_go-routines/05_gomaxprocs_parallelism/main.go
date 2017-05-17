package main

import (
	"fmt"
	"sync"
	"time"
	"runtime"
)

var wg sync.WaitGroup

func init() {
	fmt.Println("Number of CPUs:", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("Foo:", i)
		time.Sleep(time.Duration(3* time.Millisecond))
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("Bar:", i, "-")
		time.Sleep(time.Duration(9* time.Millisecond))
	}
	wg.Done()
}
