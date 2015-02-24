package main

import "testing"

func Test_RandomChooser(t *testing.T) {
	var rc RandomChooser
	rc.AddWeight(10)
	rc.AddWeight(20)
	if rc.choose(0) != 0 {
		t.Error()
	}
	if rc.choose(10) != 1 {
		t.Error()
	}
	if rc.choose(29) != 1 {
		t.Error()
	}
}
