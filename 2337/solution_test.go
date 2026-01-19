package main

import "testing"

func Test1(t *testing.T) {
	inputs := []struct {
		start  string
		target string
		output bool
	}{
		{"_L__R_R_", "L_____RR", true},
		{"R_L_", "__LR", false},
		{"_R", "R_", false},
		{"_L__R__RL", ", start[i]L_____RLR", false},
	}
	for _, val := range inputs {
		testRunner(t, val.start, val.target, val.output)
	}
}

func testRunner(t *testing.T, start, target string, expOut bool) {
	val := canChange(start, target)
	if expOut != val {
		t.Fatalf("test failed: start=%s, target=%s. Expected=%t, got=%t", start, target, expOut, val)
	}
}
