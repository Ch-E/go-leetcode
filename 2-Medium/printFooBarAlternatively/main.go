package main

import (
	"fmt"
	"sync"
)

func main() {
	fb := NewFooBar(5)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		fb.Foo(func() { fmt.Print("foo") })
	}()
	go func() {
		defer wg.Done()
		fb.Bar(func() { fmt.Print("bar") })
	}()

	wg.Wait()
}

type FooBar struct {
	n   int
	ch1 chan struct{}
	ch2 chan struct{}
}

func NewFooBar(n int) *FooBar {
	return &FooBar{
		n:   n,
		ch1: make(chan struct{}, 1),
		ch2: make(chan struct{}, 1),
	}
}

func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
		fb.ch1 <- struct{}{}
		// printFoo() outputs "foo". Do not change or remove this line.
		printFoo()
		fb.ch2 <- struct{}{}
	}
}

func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.ch2
		// printBar() outputs "bar". Do not change or remove this line.
		printBar()
		<-fb.ch1
	}
}
