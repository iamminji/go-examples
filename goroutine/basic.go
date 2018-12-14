package main

import (
	"fmt"
	"sync"
	"time"
)

// 고루틴 기본 예제

// 구조체에 해당 구조체에 채널과 waitGroup 필드 생성
type MyStruct struct {
	stop1 chan bool
	stop2 chan bool
	wait  *sync.WaitGroup
}

func NewMyStruct() *MyStruct {
	m := &MyStruct{}
	m.stop1 = make(chan bool)
	m.stop2 = make(chan bool)
	m.wait = &sync.WaitGroup{}
	return m
}

// 무한 루프 돌면서 채널 확인
func (m *MyStruct) test1() {

	defer m.wait.Done()

	for {
		select {
		case <-m.stop1:
			fmt.Println("Stop Test 1")
			return
		case <-time.After(2 * time.Second):
			fmt.Println("TEST 1 sleep 2 Seconds")
		}
	}
}

// 두번째 방법
func (m *MyStruct) test2() {

EXIT:
	for {
		select {
		case <-m.stop2:
			fmt.Println("Stop Test 2")
			break EXIT
		case <-time.After(3 * time.Second):
			fmt.Println("TEST 2 sleep 3 Seconds")
		}
	}

	m.wait.Done()
}

func main() {

	myStruct := NewMyStruct()

	myStruct.wait.Add(2)

	go myStruct.test1()
	go myStruct.test2()

	time.Sleep(3 * time.Second)
	// test1 만 종료
	myStruct.stop1 <- true

	defer myStruct.wait.Wait()
}
