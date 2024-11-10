package main
import "fmt"
var x int = 10
func average(xs[] float64) float64 {
	total := 0.0
	for _,v := range xs {
		// fmt.Println(i, v)
		total += v
	}
	return total / float64(len(xs))
}
func printName(name string){
	fmt.Println("x inside printName function : ", x)
	fmt.Println("Hello", name)
}
func multipleReturn()(int, int){
	return 5,6
}
func totalOfArray(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint){
		ret = i
		i += 2
		return
	}
}
func fact(x uint) uint {
	if x <= 1 {
		return x
	}
	return x * fact(x - 1)
}
func first(){
	fmt.Println("1st")
}
func second(){
	fmt.Println("2nd")
}
func main(){
	xs := []float64{1.2, 1.3, 1.4, 1.4}
	answer := average(xs)
	fmt.Println(answer)

	name:= "Veet"
	printName(name)
	fmt.Println("x inside main function : ", x)

	fmt.Println(multipleReturn())
	fmt.Println(totalOfArray(1,2,3,4,5,6,7,8,9,10))

	add := func(x, y int) int {
		return x + y
	}	
	fmt.Println(add(1,1))

	fmt.Println("\nEven Generator")
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())

	fmt.Println("5 factorial is : ", fact(5))
	fmt.Println("6 factorial is : ", fact(6))

	fmt.Println("Defer")
	defer second()
	first()
	first()
	defer second()


	defer func(){
		str:= recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}
