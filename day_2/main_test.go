package main

import "testing"

func TestSafeOneWrong(t *testing.T) {
	report := []int{8, 6, 4,4, 1}
    result := checkReportSafe(report)
    if result != true {
        t.Errorf("checkReportSafe returned %t; want true", result)
    }
}