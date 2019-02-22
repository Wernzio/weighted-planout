package planout

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	testCases := map[string]uint64{
		"foo":         53740260187959549,
		"":            982798738632651952,
		"foo.bar":     231603001195628059,
		"foo.bar.bar": 223288269789451194,
	}
	for key, value := range testCases {
		hashed := hash(key)
		if hashed != value {
			t.Errorf("Hash of (%s) was incorrect, got: %d, want: %d.", key, hashed, value)
		}

	}
}

func TestExecuteAllTrue(t *testing.T) {
	exp, _ := NewExp("foo", []interface{}{true, false}, []float64{1, 0})
	testCases := map[string]bool{
		"foo":     true,
		"bar":     true,
		"foo.bar": true,
		"123":     true,
		"345":     true,
		"567":     true,
		"890":     true,
		"_":       true,
		"*":       true,
		"LKJLKJ":  true,
		"LkJlKJ":  true,
	}
	for user, result := range testCases {
		decision := exp.Execute(user)
		if decision != result {
			t.Errorf("Calculated decision for UserId:(%s) with key:(%s) was incorrect, got: %t, want: %t. ", user, exp.Key, decision, result)
		}
	}
}

func TestExecuteHalfTrueHalfFalse(t *testing.T) {
	exp, _ := NewExp("foo", []interface{}{true, false}, []float64{0.5, 0.5})
	testCases := map[string]bool{
		"foo":     true,
		"bar":     true,
		"foo.bar": false,
		"123":     true,
		"345":     true,
		"567":     false,
		"890":     true,
		"_":       true,
		"*":       false,
		"LKJLKJ":  false,
		"LkJlKJ":  true,
	}
	for user, result := range testCases {
		decision := exp.Execute(user)
		if decision != result {
			t.Errorf("Calculated decision for UserId:(%s) with key:(%s) was incorrect, got: %t, want: %t. ", user, exp.Key, decision, result)
		}
	}
}

func TestExecuteAllFalse(t *testing.T) {
	exp, _ := NewExp("foo", []interface{}{true, false}, []float64{0, 1})
	testCases := map[string]bool{
		"foo":     false,
		"bar":     false,
		"foo.bar": false,
		"123":     false,
		"345":     false,
		"567":     false,
		"890":     false,
		"_":       false,
		"*":       false,
		"LKJLKJ":  false,
		"LkJlKJ":  false,
	}
	for user, result := range testCases {
		decision := exp.Execute(user)
		if decision != result {
			t.Errorf("Calculated decision for UserId:(%s) with key:(%s) was incorrect, got: %t, want: %t. ", user, exp.Key, decision, result)
		}
	}
}

func TestExecuteHandleComplexKey(t *testing.T) {
	exp, _ := NewExp("foo.bar!@#$%^&*(.!@#$%^&*(234567", []interface{}{true, false}, []float64{0.5, 0.5})
	testCases := map[string]bool{
		"foo":                      false,
		"bar":                      true,
		"foo.bar":                  true,
		"123":                      true,
		"345":                      true,
		"567":                      false,
		"890":                      true,
		"_":                        false,
		"*":                        false,
		"LKJLKJ":                   false,
		"!@#$%^&*()1234567890_+-=": false,
	}
	for user, result := range testCases {
		decision := exp.Execute(user)
		if decision != result {
			t.Errorf("Calculated decision for UserId:(%s) with key:(%s) was incorrect, got: %t, want: %t. ", user, exp.Key, decision, result)
		}
	}
}

func TestExecuteHandleUniformsDecisionWeights(t *testing.T) {
	exp, _ := NewExp("foo", []interface{}{true, false}, []float64{62, 12})
	testCases := map[string]bool{
		"foo":                      true,
		"bar":                      true,
		"foo.bar":                  false,
		"123":                      true,
		"345":                      true,
		"567":                      true,
		"890":                      true,
		"_":                        true,
		"*":                        true,
		"LKJLKJ":                   true,
		"!@#$%^&*()1234567890_+-=": true,
	}
	for user, result := range testCases {
		decision := exp.Execute(user)
		if decision != result {
			t.Errorf("Calculated decision for UserId:(%s) with key:(%s) was incorrect, got: %t, want: %t. ", user, exp.Key, decision, result)
		}
	}
}

func TestExecuteHandleMulitpleVariants(t *testing.T) {
	exp, _ := NewExp("foo", []interface{}{"true", "false", true, false, 123, 345}, []float64{1, 1, 1, 1, 1, 1})
	testCases := map[string]interface{}{
		"foo":                      "false",
		"bar":                      "false",
		"123":                      true,
		"345":                      true,
		"567":                      false,
		"890":                      true,
		"_":                        true,
		"*":                        false,
		"LKJLKJ":                   false,
		"!@#$%^&*()1234567890_+-=": false,
		"blah1":                    false,
		"blah2":                    "false",
		"blah3":                    "true",
		"blah4":                    345,
		"blah5":                    "false",
	}
	for user, result := range testCases {
		decision := exp.Execute(user)
		if decision != result {
			t.Errorf("Calculated decision for UserId:(%s) with key:(%s) was incorrect, got: %t, want: %t. ", user, exp.Key, decision, result)
		}
	}
}

func TestExecuteMismatchWeightsPercentage(t *testing.T) {
	exp, err := NewExp("foo", []interface{}{"true", "false"}, []float64{1, 1, 1})
	if exp != nil {
		t.Errorf("Expected nil as there is an error in creation")
	}
	if err == nil {
		t.Errorf("Expected error message")
	}
}

func TestUniformHash(t *testing.T) {
	exp, _ := NewExp("foo", []interface{}{true, false}, []float64{1, 1})
	testCases := map[string]float64{
		"foo":     0.4329369114744317,
		"bar":     0.4017671632807405,
		"foo.bar": 1.7702870362348528,
		"123":     0.870603965662553,
		"345":     0.8947330959080856,
		"567":     1.2553691346292328,
		"890":     0.923706126469132,
		"_":       0.8772060389443099,
		"*":       1.057521091493218,
		"LKJLKJ":  1.0539406000665854,
		"LkJlKJ":  0.4987941295076843,
	}
	for user, result := range testCases {
		hash := exp.uniformHash(user, 0.0, 2)
		if hash != result {
			fmt.Println(hash)
			fmt.Println(result)
			t.Errorf("Calculated decision for UserId:(%s) with key:(%s) was incorrect, got: %F, want: %F. ", user, exp.Key, hash, result)
		}
	}
}
