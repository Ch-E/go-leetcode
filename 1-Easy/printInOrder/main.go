package main

import (
	"fmt"
	"sync"
)

func main() {
	foo := NewFoo()
	foo.wg.Add(3)

	go foo.First(func() { fmt.Print("first") })
	go foo.Second(func() { fmt.Print("second") })
	go foo.Third(func() { fmt.Print("third") })

	foo.ch1 <- struct{}{}
	foo.wg.Wait()
}

type Foo struct {
	ch1 chan struct{}
	ch2 chan struct{}
	ch3 chan struct{}
	wg  sync.WaitGroup
}

func NewFoo() *Foo {
	return &Foo{
		ch1: make(chan struct{}, 1),
		ch2: make(chan struct{}),
		ch3: make(chan struct{}),
	}
}

func (f *Foo) First(printFirst func()) {
	defer f.wg.Done()
	<-f.ch1 // wait for something in channel 1
	printFirst()
	f.ch2 <- struct{}{} // send something into channel 2 to trigger it
}

func (f *Foo) Second(printSecond func()) {
	defer f.wg.Done()
	<-f.ch2
	printSecond()
	f.ch3 <- struct{}{}
}

func (f *Foo) Third(printThird func()) {
	defer f.wg.Done()
	<-f.ch3
	printThird()
}
