package main

import ("fmt"; "math")

type Shape interface {
	perimeter() float64
}
func (c* Cicle) perimeter() float64 {
	return 2 * math.Pi * c.r
}
func (r* Ractangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return 2 * (l + w)
}
func main() {
	xs := [5]int{1, 2, 5, 4, 3}
	fmt.Println(maxVal(xs))
}
