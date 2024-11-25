package main

import (
	"fmt"
	"sync"
	//"time"
)

func main() {
	//task1()
	//task2()
	//task3()
	//task4()
	task5()

	//time.Sleep(2 * time.Second)
}

func task1() {

	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Printf("%d\n", i)
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%c\n", 'a'+i)
		}
	}()

}

func task2() {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			fmt.Printf("%d\n", i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Printf("%c\n", 'a'+i)
		}
	}()

	wg.Wait()
}

func task3() {
	c := make(chan int)
	done := make(chan bool)

	go producer(c)
	go consumer(c, done)

	// Wait for consumer to finish
	<-done
}

func producer(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

func consumer(ch chan int, done chan bool) {
	for num := range ch {
		fmt.Println("", num)
	}
	done <- true
}

func task4() {
	c := make(chan int, 5)
	done := make(chan bool)

	go producer(c)
	go consumer(c, done)

	// Wait for consumer to finish
	<-done

	//buffered channel only send data when the buffer is full, while unbuffered just send data immediately
}

func task5() {
	cEven := make(chan int)
	cOdd := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 1; i <= 20; i++ {
			if i%2 == 0 {
				cEven <- i
			} else {
				cOdd <- i
			}
		}
		close(cEven)
		close(cOdd)
	}()

	go func() {
		for {
			select {
			case even, ok := <-cEven:
				if ok {
					fmt.Println("Received an Even Number:", even)
				}
			case odd, ok := <-cOdd:
				if ok {
					fmt.Println("Received an Odd Number:", odd)
				}
			}

		}
	}()

	<-done
}
