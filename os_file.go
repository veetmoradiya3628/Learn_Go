package main
import (
	"fmt";
	"os";
)
func main(){
	file, err := os.Open("test.txt")
	if err != nil {
		return
	}

	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return
	}
	fmt.Println("stats: ", stat)
	fmt.Println("stats size: ", stat.Size())
	fmt.Println("stats name: ", stat.Name())
	fmt.Println("stats mode: ", stat.Mode())
	
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println(str)

	file, err = os.Create("test2.txt")
	if err != nil {
		return
	}
	defer file.Close()
	fmt.Println(file)

	errDelete := os.Remove("test2.txt")
	if errDelete != nil {
		fmt.Println("Error deleting file", errDelete)
		return
	}
}
