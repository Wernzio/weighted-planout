package planout

import (
	"crypto/sha1"
	"fmt"
	"strconv"
)

type Experiment struct {
	Key         string
	UserId      string
	Choices     []interface{}
	Percentages []float64
	sum         float64
	cweights    []float64
}

func hash(in string) uint64 {
	var x [20]byte = sha1.Sum([]byte(in))
	var y string = fmt.Sprintf("%x", x[0:8])
	y = y[0 : len(y)-1]
	var z uint64 = 0
	z, _ = strconv.ParseUint(y, 16, 64)
	return z
}

func NewExp(key string, choices []interface{}, percentages []float64) (*Experiment, error) {
	if len(choices) != len(percentages) {
		return nil, fmt.Errorf("Percentage and weights must match")
	}
	exp := Experiment{Key: key, Choices: choices, Percentages: percentages}
	exp.acummulativeWeights()
	return &exp, nil
}

func (exp *Experiment) Execute(UserId string) interface{} {
	stop_val := exp.uniformHash(UserId, 0.0, exp.sum)
	var decision interface{}
	for i := range exp.cweights {
		if stop_val <= exp.cweights[i] {
			decision = exp.Choices[i]
			break
		}
	}
	return decision
}

func (exp *Experiment) uniformHash(UserId string, min, max float64) float64 {
	scale, _ := strconv.ParseUint("FFFFFFFFFFFFFFF", 16, 64)
	h := hash(exp.Key + "." + UserId)
	shift := float64(h) / float64(scale)
	return min + shift*(max-min)
}

func (exp *Experiment) acummulativeWeights() {
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
