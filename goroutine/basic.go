package main

import (
	"fmt"
	"sync"
	"time"
)

// 고루틴 기본 예제

// 구조체에 해당 구조체에 채널과 waitGroup 필드 생성
type MyStruct struct {
	stop chan bool
	wait *sync.WaitGroup
}

func NewMyStruct() *MyStruct {
	m := &MyStruct{}
	m.stop = make(chan bool)
	m.wait = &sync.WaitGroup{}
	return m
}

// 무한 루프 돌면서 채널 확인
func (m *MyStruct) test() {

	defer m.wait.Done()

	for {
		select {
		case <-m.stop:
			fmt.Println("Stop Test")
			// return 혹은 break 시에는 for 루프에 라벨을 붙여서 break 해야 한다.
			return
		case <-time.After(2 * time.Second):
			fmt.Println("sleep 2 Seconds")
		}
	}
}

func main() {

	myStruct := NewMyStruct()
	myStruct.wait.Add(1)
	go myStruct.test()

	// 종료 시 아래 주석 참고
	// myStruct.stop <- true

	defer myStruct.wait.Wait()
}
