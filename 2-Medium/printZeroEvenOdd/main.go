/*
You have a function printNumber that can be called with an integer parameter and prints it to the console.

For example, calling printNumber(7) prints 7 to the console.
You are given an instance of the class ZeroEvenOdd that has three functions: zero, even, and odd.
The same instance of ZeroEvenOdd will be passed to three different threads:

Thread A: calls zero() that should only output 0's.
Thread B: calls even() that should only output even numbers.
Thread C: calls odd() that should only output odd numbers.
Modify the given class to output the series "010203040506..." where the length of the series must be 2n.

Implement the ZeroEvenOdd class:

ZeroEvenOdd(int n) Initializes the object with the number n that represents the numbers that should be printed.
void zero(printNumber) Calls printNumber to output one zero.
void even(printNumber) Calls printNumber to output one even number.
void odd(printNumber) Calls printNumber to output one odd number.
*/

package main

import "fmt"

type ZeroEvenOdd struct {
	n      int
	chZero chan struct{}
	chEven chan struct{}
	chOdd  chan struct{}
}

func Constructor(n int) ZeroEvenOdd {
	return ZeroEvenOdd{
		n:      n,
		chZero: make(chan struct{}, 1),
		chEven: make(chan struct{}, 1),
		chOdd:  make(chan struct{}, 1),
	}
}

func (z *ZeroEvenOdd) Zero(printNumber func(int)) {
	for i := 1; i <= z.n; i++ {
		<-z.chZero
		printNumber(0)
		if i%2 == 1 {
			z.chOdd <- struct{}{}
		} else {
			z.chEven <- struct{}{}
		}
	}
}

func (z *ZeroEvenOdd) Even(printNumber func(int)) {

	for i := 2; i <= z.n; i += 2 {
		<-z.chEven
		printNumber(i)
		z.chZero <- struct{}{}
	}
}

func (z *ZeroEvenOdd) Odd(printNumber func(int)) {
	for i := 1; i <= z.n; i += 2 {
		<-z.chOdd
		printNumber(i)
		z.chZero <- struct{}{}
	}
}

func main() {
	printNumber := func(x int) {
		fmt.Print(x)
	}

	obj := Constructor(5)
	obj.chZero <- struct{}{}
	go obj.Zero(printNumber)
	go obj.Even(printNumber)
	go obj.Odd(printNumber)

	// Wait for goroutines to finish
	select {}
}
