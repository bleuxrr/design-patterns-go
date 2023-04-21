package main

import "fmt"

type lfu struct {
}

func (l *lfu) evict(c *cache) {
	fmt.Println("evicting by lfu strategy")
}
