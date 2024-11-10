package mathutil

// Avg calculates the average of a slice of float64 numbers.
// If the slice is empty, it returns 0.
func Avg(numbers []float64) float64 {
    if len(numbers) == 0 {
        return 0
    }

    sum := 0.0
    for _, num := range numbers {
        sum += num
    }
    return sum / float64(len(numbers))
}
