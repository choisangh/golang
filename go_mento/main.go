package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func nextValue() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	runtime.GOMAXPROCS(6)
	start := time.Now()

	var sum int64
	var wg sync.WaitGroup
	var mu sync.Mutex
	ch := make(chan int)
	count := 0
	n := 5

	wg.Wait()

	for i := 0; i < 1000000000; i += 1000000000 / n {
		wg.Add(1)

		go sumation(i, i+1000000000/n, &sum, &mu, &wg, ch, &count, n)
	}

	for receive := range ch {
		fmt.Printf("%d thread 50%% done! \n", receive)
	}

	wg.Wait()

	elapsed := time.Since(start)

	fmt.Println(sum, elapsed)
}

func sumation(start, end int, sumTotal *int64, mu *sync.Mutex, wg *sync.WaitGroup, ch chan int, count *int, n int) {
	var sum int64
	for i := start; i < end/2; i++ {
		sum += int64(i)
	}

	ch <- start / 200000000

	for i := end / 2; i < end; i++ {
		sum += int64(i)
	}

	mu.Lock()
	*sumTotal += sum
	*count += 1
	mu.Unlock()
	if *count == n {
		close(ch)
	}
	wg.Done()

}
