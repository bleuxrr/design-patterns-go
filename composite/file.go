package main

import "fmt"

type file struct {
	name string
}

func (f *file) search(keyword string) {
	fmt.Printf("searching for keyword %s in file %s\n", keyword, f.name)
}
