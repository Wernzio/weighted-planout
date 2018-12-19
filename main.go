package main

import (
	"fmt"
	planout "github.com/Wernzio/weighted-planout"
)

func main() {
	nexExp := planout.Experiment{Key: "blah", PID: "blah", Choices: []interface{}{true, false}, Percentages: []float64{0.5, 0.5}}
	fmt.Println(nexExp.Execute())
}
