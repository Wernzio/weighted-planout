package planout

import (
	"crypto/sha1"
	"fmt"
	"strconv"
)

type planoutParams struct {
	Choice     string
	Percentage string
}

type Experiment struct {
	Key         string
	PID         string
	Choices     []interface{}
	Percentages []float64
}

func hash(in string) uint64 {
	// Compute 20- byte sha1
	var x [20]byte = sha1.Sum([]byte(in))
	// Get the first 15 characters of the hexdigest.
	var y string = fmt.Sprintf("%x", x[0:8])
	y = y[0 : len(y)-1]
	// Convert hex string into uint64
	var z uint64 = 0
	z, _ = strconv.ParseUint(y, 16, 64)
	return z
}

func (exp *Experiment) getUniform(min, max float64) float64 {
	scale, _ := strconv.ParseUint("FFFFFFFFFFFFFFF", 16, 64)
	h := exp.getHash()
	shift := float64(h) / float64(scale)
	return min + shift*(max-min)
}

func (exp *Experiment) getHash() uint64 {
	return hash(exp.Key + "." + exp.PID)
}

func (exp *Experiment) getCummulativeWeights() (float64, []float64) {
	nweights := len(exp.Percentages)
	cweights := make([]float64, nweights)
	sum := 0.0
	for i := range exp.Percentages {
		sum = sum + exp.Percentages[i]
		cweights[i] = sum
	}
	return sum, cweights
}

func (exp *Experiment) Execute() interface{} {
	sum, cweights := exp.getCummulativeWeights()
	stop_val := exp.getUniform(0.0, sum)
	fmt.Println(stop_val)
	for i := range cweights {
		if stop_val <= cweights[i] {
			return exp.Choices[i]
		}
	}
	return nil
}
