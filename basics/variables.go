package main
import "fmt"

func main(){
	var x string = "Hello World"
	fmt.Println(x)

	var str1 string = "Hello"
	var str2 string = "World"
	fmt.Println(str1 == str2)
	fmt.Println(str1 != str2)

	var num1 = 10
	fmt.Println(num1)

	dogName := "Max"
	fmt.Println(`My dog's name is ` + dogName)

	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println(output)
}
