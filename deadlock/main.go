package main

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	mu    sync.Mutex
	value int
}

func main() {
	var wg sync.WaitGroup
	printSum := func(v1, v2 *value) {
		defer wg.Done()
		v1.mu.Lock()
		defer v1.mu.Unlock()

		time.Sleep(2 * time.Second)
		v2.mu.Lock()
		defer v2.mu.Unlock()

		fmt.Printf("sum=&v\n", v1.value+v2.value)
	}

	var a, b value
	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()

	// 相互排他：printSum関数はaとbの両方に対して排他的アクセス権を必要としている
	// 条件待ち：printSumはaもしくはbのどちらかを保持していて、もう片方を持っている
	// 横取り不可：ゴルーチンを横取りする方法は提供されていない
	// 循環待ち：printSumの最初の呼び出しでは2番目の呼び出しを待っていて、逆もまた然り
	// fatal error: all goroutines are asleep - deadlock!
}
