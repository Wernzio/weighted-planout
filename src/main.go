package main

import (
	planout "./goPlanout"
	"fmt"
)

func main() {
	nexExp := planout.Experiment{Key: "blah", PID: "blah", Choices: []interface{}{true, false}, Percentages: []float64{0.5, 0.5}}
	fmt.Println(nexExp.Execute())
}
