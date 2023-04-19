package main

import (
	"fmt"
	"sync"
)

var once sync.Once

type single struct{}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("creating single instance now")
				singleInstance = &single{}
			})
	} else {
		fmt.Println("single instance already created-2")
	}
	return singleInstance
}
