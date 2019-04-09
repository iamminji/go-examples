package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4
	taskLoad         = 10
)

var wg2 sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func worker(tasks chan string, worker int) {
	defer wg2.Done()

	for {
		task, ok := <-tasks
		if !ok {
			fmt.Printf("작업자: %d : 종료합니다.\n", worker)
			return
		}
		fmt.Printf("작업자: %d : 작업 시작: %s\n", worker, task)

		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		fmt.Printf("작업자: %d : 작업 완료: %s\n", worker, task)
	}
}

func main() {
	tasks := make(chan string, taskLoad)

	wg2.Add(numberGoroutines)
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("작업 :%d", post)
	}

	close(tasks)
	wg2.Wait()
}
