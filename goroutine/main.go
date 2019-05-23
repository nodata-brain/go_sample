package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		index := i
		wg.Add(1) //wgのインクリメントはgo funcの外に書く
		go func() {
			defer wg.Done() //完了したらwgをデクリメント
			fmt.Println("success!", index)
		}()
	}
	wg.Wait() //goroutineが終了するまでブロックする

}
