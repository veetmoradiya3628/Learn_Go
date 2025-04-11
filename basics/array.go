package main
import "fmt"

func main(){
	// array
	var x[5] int
	x[4] = 10
	x[1] = 5
	fmt.Println(x)

	// array with values and avg calculation
	var arr [5]float64
	arr[0] = 98
	arr[1] = 93
	arr[2] = 77
	arr[3] = 82
	arr[4] = 83
	var total float64 = 0
	for i := 0; i < 5; i++ {
		total += arr[i]
	}
	fmt.Println(total / 5)

	// array with values and avg calculation using range
	total = 0
	for _, value := range x {
		total += float64(value)
	}
	fmt.Println(total / float64(len(x)))
}
