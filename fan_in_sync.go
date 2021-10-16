package main

import "fmt"

// chan are first class variables that means we can have chan within chan. below is the pattern to sync the async output of the fan
// in pattern. with the message in the chan it also has a waitforit chan bool which block while listning on it until the reader of the
// message (main goroutine) has read and then is read to send the true signal back to the boring func to print out more stuff
type Message struct {
	msg  string
	wait chan bool
}

func fan_sync_driver() {
	ck := fan(boring("amy"), boring("jon"))
	for i := 0; i < 5; i++ {
		msg1 := <-ck
		fmt.Println(msg1.msg)
		msg2 := <-ck
		fmt.Println(msg2.msg)
		msg1.wait <- true
		msg2.wait <- true
	}
}

func fan(c, k <-chan Message) <-chan Message {
	ck := make(chan Message)
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

func boring(id string) <-chan Message {
	waitForIt := make(chan bool)
	mssg := make(chan Message)
	go func() {
		for i := 0; ; i++ {
			mssg <- Message{fmt.Sprintf("%s: %d", id, i), waitForIt}
			<-waitForIt
		}
	}()
	return mssg
}
