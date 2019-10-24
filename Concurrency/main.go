package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//concurrency0()
	//concurrency1()
	//concurrency2()
	concurrency3()
}

func concurrency0() {
	go countinfinte("sheep")
	countinfinte("dog")
}

func concurrency2() {
	c := make(chan string)
	go countChannel("sheep", c)
	for {
		msg, open := <-c
		//check if channel is closed!
		if !open {
			break
		}
		fmt.Println(msg)
	}

}

func concurrency3() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep((time.Millisecond * 500))
		}
	}()

	go func() {
		for {
			c2 <- "Every two seconds"
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		fmt.Println(<-c1)
		fmt.Println(<-c2)
	}
}

func countChannel(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}
	//NEVER FORGET TO CLOSE THE CHANNEL in the used function
	close(c)
}

func concurrency1() {
	//WaitGroup, code below go func will only continue
	//if all go functions inside have finished!
	var wg sync.WaitGroup
	//Two go routines
	wg.Add(2)
	go func() {
		count("sheep", 5)
		wg.Done()
	}()
	go func() {
		count("dog", 10)
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("More Code to come if everything is done")
}

func count(thing string, idx int) {
	for i := 1; i <= idx; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}

func countinfinte(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
