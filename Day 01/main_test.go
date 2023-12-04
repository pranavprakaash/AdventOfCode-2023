package main

import "testing"

func TestCalculateCalibrationValue(t *testing.T) {
	result := calculateCalibrationValue("1somethingtwo3xeonfour7nine")
	if result != 17 {
		t.Errorf("Result was incorrect, got: %d, want:%d.", result, 17)
	}
}
