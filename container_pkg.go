package main
import (
	"fmt";
	"container/list";
	"sort"
)
type Person struct {
	Name string
	Age int
}
type ByName []Person
func (ps ByName) Len() int { return len(ps) }
func (ps ByName) Less(i, j int) bool {
	return ps[i].Name < ps[j].Name
}
func (ps ByName) Swap(i, j int){
	ps[i], ps[j] = ps[j], ps[i]
}
func main(){
	var l list.List
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}

	kids := []Person{
		{"Jill", 9},
		{"Jack", 10},
	}
	sort.Sort(ByName(kids))
	fmt.Println(kids)
}
