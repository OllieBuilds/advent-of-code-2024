package main

import "testing"

func TestGetDistLeft(t *testing.T) {
    result := getDist(2, 3)
    if result != 1 {
        t.Errorf("getDist(2, 3) = %d; want 1", result)
    }
}
func TestGetDistRight(t *testing.T) {
    result := getDist(8, 1)
    if result != 7 {
        t.Errorf("getDist(8, 1) = %d; want 7", result)
    }
}
func TestGetDistEqual(t *testing.T) {
    result := getDist(3, 3)
    if result != 0 {
        t.Errorf("getDist(3, 3) = %d; want 3", result)
    }
}

