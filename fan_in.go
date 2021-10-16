package main

import "fmt"

// this is the fan in pattern where two or more chan combines to give one single chan.
func fanindriver() {
	ck := fanin(counts("amy"), counts("rose"))
	// amychan := counts("amy")
	// rosechan := counts("rose")

	for i := 0; i < 10; i++ {
		fmt.Println(<-ck)
	}
}

func fanin(c, k <-chan string) <-chan string {
	ck := make(chan string)
	go func() {
		for {
			ck <- <-c
		}
	}()
	go func() {
		for {
			ck <- <-k
		}
	}()
	return ck
}


/*
rose: 0
rose: 1
amy: 0
amy: 1
amy: 2
amy: 3
rose: 2
amy: 4
rose: 3
rose: 4
*/
