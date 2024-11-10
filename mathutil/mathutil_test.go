package mathutil

import (
    "testing"
)

// TestAvg tests the Avg function.
func TestAvg(t *testing.T) {
    // Define test cases
    tests := []struct {
        name     string
        input    []float64
        expected float64
    }{
        {"average of positive numbers", []float64{1, 2, 3, 4, 5}, 3},
        {"average of negative numbers", []float64{-1, -2, -3, -4, -5}, -3},
        {"mixed positive and negative", []float64{1, -1, 2, -2, 3, -3}, 0},
        {"single number", []float64{5}, 5},
        {"empty slice", []float64{}, 0},
    }

    // Execute each test case
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Avg(tt.input)
            if result != tt.expected {
                t.Errorf("Avg(%v) = %v; expected %v", tt.input, result, tt.expected)
            }
        })
    }
}
