package main
import "fmt"

func main(){
	// slice
	x := make([]float64, 5)
	fmt.Println(x)

	arr := [5]float64{1,2,3,4,5}
	y := arr[0:2]
	fmt.Println(y)

	// append
	slice1 := []int{1,2,3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	// copy
	slice3 := []int{1,2,3}
	slice4 := make([]int, 2)
	copy(slice4, slice3)
	fmt.Println(slice4, slice3)
}
