package main

import "fmt"

func main() {
	x, y := 10,20
	
	p := &x
	q := &y

	fmt.Println(p)
	*p = 30
	*q = 10
	fmt.Println(*p)
	
	//pass by val
	// swap(x,y)
	// fmt.Println(x,y)

	//pass by ref
	swap(&x,&y) // or swap(p,q)
	fmt.Println(x,y)
}

//pass by value
// func swap(x,y int){
	// temp := x
	// x = y
	// y = temp
// }

//pass by ref
func swap(x,y *int){
	temp := *x
	*x = *y
	*y = temp
}
