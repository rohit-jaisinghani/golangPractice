package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)


func sum_go(a, b float64, wg *sync.WaitGroup) {
	s := a + b
	fmt.Printf("Sum of %.2f and %.2f is %.2f\n", a, b, s)
	wg.Done()
}

func f1(wg *sync.WaitGroup) {
	fmt.Println("f1() execution started")
	for i := 0; i < 3; i++ {
		fmt.Println("f1, i=", i)
		// sleep for a second to simulate an expensive task.
		time.Sleep(time.Second)

	}
	fmt.Println("f1 execution finished")
	wg.Done()
}

func f2(wg *sync.WaitGroup) {
	fmt.Println("f2() execution started")
	time.Sleep(time.Second)

	for i := 5; i < 8; i++ {
		fmt.Println("f2(), i=", i)

	}
	fmt.Println("f2 execution finished")
	wg.Done()

}
func f3() {
	fmt.Println("f3() execution started")
	time.Sleep(time.Second)

	for i := 5; i < 8; i++ {
		fmt.Println("f3(), i=", i)

	}
	fmt.Println("f3 execution finished")

}
func main() {
	fmt.Println("execution started")
	var wg sync.WaitGroup
	wg.Add(5)
	go f2(&wg)
	go f1(&wg)

	go sum_go(10.3, 20.1, &wg)
	go sum_go(5, 7, &wg)
	go sum_go(33.5, 33.66, &wg)

	f3()

	fmt.Println("no of goroutines", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("execution stopped")

}
