package planout

import (
	"crypto/sha1"
	"fmt"
	"strconv"
)

type Experiment struct {
	Key         string
	PID         string
	Choices     []interface{}
	Percentages []float64
	sum         float64
	cweights    []float64
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

func NewExp(key string, choices []interface{}, percentages []float64) *Experiment {
	exp := Experiment{Key: key, Choices: choices, Percentages: percentages}
	exp.getCummulativeWeights()
	return &exp
}

func (exp *Experiment) Execute(PID string) interface{} {
	stop_val := exp.getUniform(PID, 0.0, exp.sum)
	for i := range exp.cweights {
		if stop_val <= exp.cweights[i] {
			return exp.Choices[i]
		}
	}
	return nil
}

func (exp *Experiment) getUniform(PID string, min, max float64) float64 {
	scale, _ := strconv.ParseUint("FFFFFFFFFFFFFFF", 16, 64)
	h := hash(exp.Key + "." + PID)
	shift := float64(h) / float64(scale)
	return min + shift*(max-min)
}

func (exp *Experiment) getCummulativeWeights() {
	nweights := len(exp.Percentages)
	cweights := make([]float64, nweights)
	sum := 0.0
	for i := range exp.Percentages {
		sum = sum + exp.Percentages[i]
		cweights[i] = sum
	}
	exp.sum = sum
	exp.cweights = cweights
}
