package main

import "fmt"

type fifo struct {
}

func (l *fifo) evict(c *cache) {
	fmt.Println("evicting by fifo strategy")
}
