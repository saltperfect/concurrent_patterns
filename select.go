package main

import (
	"fmt"
	"time"
)

func selectdriver() {

	c := counts("rose")

	for i := 0; i < 20; i++ {
		select {
		case r := <-c:
			fmt.Println(r)
		case <-time.After(time.Second):
			fmt.Println("too slow")
		}
	}
}

func mainTimerDriver() {
	c := counts("rose")

	tc := time.After(time.Second * 5)

	for {
		select {
		case r := <-c:
			fmt.Println(r)
		case <-tc:
			fmt.Println("you talk too  much")
			return
		}
	}
}
