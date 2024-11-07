package main
import "fmt"

func main(){
	fmt.Println("first for loop")
	// for loop
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	fmt.Println("2nd for loop")
	for j := 1; j <= 10; j++ {
		fmt.Println(j)
	}

	// if else
	fmt.Println("if else statement")
	for i := 1; i <= 10; i++ {
		if i % 2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}

	// switch case
	fmt.Println("switch case")
	for i := 1; i <= 10; i++ {
		switch i {
		case 1: fmt.Println(i, "one")
		case 2: fmt.Println(i, "two")
		case 3: fmt.Println(i, "three")
		case 4: fmt.Println(i, "four")
		case 5: fmt.Println(i, "five")
		case 6: fmt.Println(i, "six")
		case 7: fmt.Println(i, "seven")
		case 8: fmt.Println(i, "eight")
		case 9: fmt.Println(i, "nine")
		case 10: fmt.Println(i, "ten")
		default: fmt.Println(i, "unknown number")
		}
	}
}
