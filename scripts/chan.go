package main

import (
	"fmt"
	"sync"
	"time"
)

const MAX = 5

func worker(id int, wk *sync.WaitGroup, out chan int) {
	fmt.Println("Sleeping worker ", id)
	time.Sleep(time.Duration(id * int(time.Second)))
	fmt.Println("Sending id to channel ", id)
	out <- id
	fmt.Println("Done worker ", id)

	wk.Done()
}

func sendOneByOne(wg *sync.WaitGroup, out chan int) {
	out <- 1
	fmt.Println("Done Sending ", 1)
	out <- 2
	fmt.Println("Done Sending ", 2)
	out <- 3
	fmt.Println("Done Sending ", 3)
	out <- 4
	fmt.Println("Done Sending ", 4)
	out <- 5
	fmt.Println("Done Sending ", 4)

	fmt.Println("Sent all")
	wg.Done()
}

func main() {
	out := make(chan int, 1)
	var wg sync.WaitGroup

	// for i := 1; i <= MAX; i++ {
	// 	wg.Add(1)
	// 	go worker(i, &wg, out)
	// }
	wg.Add(1)
	go sendOneByOne(&wg, out)

	fmt.Println("Sleeping before out 1")
	time.Sleep(5 * time.Second)
	out1 := <-out
	fmt.Println("Out 1 ", out1, "Sleeping")
	time.Sleep(5 * time.Second)
	out2 := <-out
	fmt.Println("Out 2 ", out2, "Sleeping")
	time.Sleep(5 * time.Second)
	out3 := <-out
	fmt.Println("Out 3 ", out3, "Sleeping")
	time.Sleep(5 * time.Second)
	out4 := <-out
	fmt.Println("Out 4 ", out4, "Sleeping")
	time.Sleep(5 * time.Second)
	out5 := <-out
	fmt.Println("Out 5 ", out5, "Sleeping")
	time.Sleep(5 * time.Second)

	wg.Wait()
}
