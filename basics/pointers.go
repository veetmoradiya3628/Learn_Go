package main
import "fmt"
func zero(xPtr *int){
	*xPtr = 0
}
func main(){
	x := 5
	zero(&x)
	fmt.Println(x)

	xPtr := new(int)
	zero(xPtr)
	fmt.Println(*xPtr)
}
