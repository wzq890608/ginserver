package mytest

import (
	"sync"
	"fmt"
	"time"
)

func TestThread1(wg sync.WaitGroup)  {
	for i:=0;i<3;i++{
		fmt.Println("one---",i)
		time.Sleep(2*time.Second)
	}
	wg.Done()
}

func TestThread2(wg sync.WaitGroup)  {
	for i:=0;i<5;i++{
		fmt.Println("two---",i)
		time.Sleep(4*time.Second)
	}
	wg.Done()
}