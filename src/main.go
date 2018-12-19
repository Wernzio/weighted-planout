package main

import (
	planout "./goPlanout"
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	nexExp := planout.Experiment{Key: "blah", Choices: []interface{}{true, false}, Percentages: []float64{0.5, 0.5}}
	nexExp.Init()
	var wg sync.WaitGroup
	i := 0
	for i < 1000000 {
		wg.Add(1)
		go nexExp.Execute("blah", &wg)
		i++
	}
	wg.Wait()
	fmt.Println(time.Since(start))
}
