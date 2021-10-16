package main

func main() {
	mainTimerDriver()
}
// func main() {
// 	ck := fanin(counts("amy"), counts("rose"))
// 	// amychan := counts("amy")
// 	// rosechan := counts("rose")

// 	for i := 0; i < 10; i++ {
// 		fmt.Println(<-ck)
// 	}
// }

// func fanin(c, k <-chan string) <-chan string {
// 	ck := make(chan string)
// 	go func() {
// 		for{ ck <- <- c}
// 	}()
// 	go func() {
// 		for{ ck <- <- k}
// 	}()
// 	return ck
// }

// func counts(name string) <-chan string {
// 	c := make(chan string)
// 	go func() {
// 		for i := 0; ; i++ {
// 			c <- fmt.Sprintf("%s: %d", name, i)
// 		}
// 	}()
// 	return c
// }
