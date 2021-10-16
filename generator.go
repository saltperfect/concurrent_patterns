package main

import (
	"fmt"
	"math/rand"
	"time"
)

// this is a generator pattern where if you want to share data from one goroutine to another, we can have a function return a chan and
// the reciever of the chan can listen to it.
func gendriver() {
	amychan := counts("amy")
	rosechan := counts("rose")

	for i := 0; i < 5; i++ {
		fmt.Println(<-amychan)
		fmt.Println(<-rosechan)
	}
}

func counts(name string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s: %d", name, i)
			time.Sleep(time.Duration(rand.Int31n(2e5) * int32(time.Millisecond)))
		}
	}()
	return c
}

/*
Outputs:
amy: 0
rose: 0
amy: 1
rose: 1
amy: 2
rose: 2
amy: 3
rose: 3
amy: 4
rose: 4
*/
