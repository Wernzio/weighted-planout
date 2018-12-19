package main

import (
	planout "./goPlanout"
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	nexExp := planout.NewExp("blah", []interface{}{true, false}, []float64{0.05, 0.95})
	var wg sync.WaitGroup
	i := 0
	for i < 1000000 {
		wg.Add(1)
		go nexExp.Execute("blah")
		i++
	}
	wg.Wait()
	fmt.Println(time.Since(start))
}
