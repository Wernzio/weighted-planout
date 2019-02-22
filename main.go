package main

import (
	planout "./planout"
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	nexExp, _ := planout.NewExp("blah", []interface{}{true, false}, []float64{0.05, 0.95})
	var wg sync.WaitGroup
	i := 0
	for i < 10000000 {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			nexExp.Execute("blah")
			wg.Done()
		}(&wg)
		i++
	}
	wg.Wait()
	fmt.Println(time.Since(start))
}
