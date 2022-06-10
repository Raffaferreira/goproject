package goroutines

import (
	"fmt"
	"time"
)

func MainConcurrencyGoRoutines() {
	go count("Sheep")
	go count("Fish")

	fmt.Scanln()
}

func MainSecondGoRoutines() {
	c := make(chan string)
	go countChannel("Sheep", c)

	// for {
	// 	mensagem, open := <-c
	// 	if !open {
	// 		break
	// 	}
	// 	fmt.Println(mensagem)
	// }

	for mensagem := range c {
		fmt.Println(mensagem)
	}

}

func MainThirdGoRoutines() {
	c := make(chan string, 2)
	c <- "hello"
	c <- "world"
	c <- "world" //deadlock

	msg := <-c
	fmt.Println(msg)

	msg = <-c
	fmt.Println(msg)

}

func MainForthGoRoutines() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "Every 2 Seconds"
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}

}

func MainFifthGoRoutines() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	go workerfib(jobs, results)
	go workerfib(jobs, results)
	go workerfib(jobs, results)
	go workerfib(jobs, results)
	go workerfib(jobs, results)

	for i := 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 0; i < 100; i++ {
		fmt.Println(<-results)
	}
}

func count(thing string) {
	for i := 1; i <= 10; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}

func countChannel(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}

	close(c)
}

func workerfib(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n+2)
}
