package main

import "fmt"

type A struct {
	v1 int
	v2 int
}

type Option func(*A)

func V1(v1 int) Option {
	return func(a *A) {
		a.v1 = v1
	}
}

func V2(v2 int) Option {
	return func(a *A) {
		a.v2 = v2
	}
}

func NewA(opts ...Option) *A {
	a := &A{}
	for _, opt := range opts {
		opt(a)
	}
	return a
}

func main() {
	fmt.Printf("%#v\n", NewA())               // &A{v1:0, v2:0}
	fmt.Printf("%#v\n", NewA(V1(42)))         // &A{v1:42, v2:0}
	fmt.Printf("%#v\n", NewA(V2(42)))         // &A{v1:0, v2:0}
	fmt.Printf("%#v\n", NewA(V1(42), V2(42))) // &A{v1:42, v2:0}
}
