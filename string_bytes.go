package main
import (
	"fmt";
	"strings"
	"bytes"
)
func main(){
	fmt.Println(strings.Contains("test", "es"))
	fmt.Println(strings.Count("test", "t"))
	fmt.Println(strings.HasPrefix("test", "te"))
	fmt.Println(strings.HasSuffix("test", "t"))
	fmt.Println(strings.HasSuffix("test", "mest"))
	fmt.Println(strings.Index("test", "e"))
	fmt.Println(strings.Index("test", "m"))
	fmt.Println(strings.Join([]string{"a", "b", "c", "d"}, "-"))
	fmt.Println(strings.Repeat("#", 10))
	fmt.Println(strings.Replace("aaaa", "a", "b", 2))
	fmt.Println(strings.Split("a-b-c-d-e", "-"))
	var splittedStr = strings.Split("a-b-c-d-e", "-")
	for _, val := range splittedStr {
		fmt.Println(val + " ")
	}

	fmt.Println(strings.ToLower("TEST"))
	fmt.Println(strings.ToUpper("test"))
	fmt.Println(strings.TrimSpace(" test "))
	fmt.Println(strings.Trim(" test ", " "))
	fmt.Println(strings.TrimLeft(" test ", " "))
	fmt.Println(strings.TrimRight(" test ", " "))
	fmt.Println(strings.TrimPrefix("test", "te"))
	fmt.Println(strings.TrimSuffix("test", "st"))
	fmt.Println(strings.Fields("test test test"))

	arr := []byte("test")
	fmt.Println(arr)
	fmt.Println(string(arr))

	var buf bytes.Buffer
	buf.Write([]byte("test"))
}
