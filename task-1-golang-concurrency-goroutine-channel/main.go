package main

import (
	"fmt"
	"sync"
)

func worker(n int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- n * 2
}

func main() {
	numbers := []int{3, 6, 9, 11, 14}
	fmt.Println("Angka awal:", numbers)

	ch := make(chan int, len(numbers))
	var wg sync.WaitGroup

	for _, n := range numbers {
		wg.Add(1)
		go worker(n, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	var results []int
	for res := range ch {
		results = append(results, res)
	}

	fmt.Println("Hasil proses:", results)
}
