package main
import "fmt"

func main(){

	// Array declarations
	var arr1 [3]int                // array of 3 ints, zero-initialized
	arr2 := [3]int{1, 2, 3}        // array with explicit length and values
	arr3 := [...]int{4, 5, 6, 7}   // array with inferred length
	// size dene ki zarurat nahi hai
	var arr4 = [5]string{"a", "b"} // array of 5 strings, rest are empty

	//question: can we add elements to an array?
	// No, you cannot add elements to arr1 because arrays in Go have a fixed size.
	// To "add" elements, use a slice instead.

	//but if i define it like this
	//arr5 := make([]int, 0) 
	// or
	//arr5 := []int{}
	// then i can add elements to it, because slices are dynamic

	fmt.Println("Arrays:")
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)

	// Slice declarations
	var s1 []int                   // nil slice
	s2 := []int{10, 20, 30}        // slice with values
	s3 := make([]int, 4)           // slice of length 4, zero-initialized
	s4 := arr2[:]                  // slice from array (full)
	s5 := arr3[1:3]                // slice from array (partial)
	s6 := s2[1:]                   // slice from slice

	fmt.Println("Slices:")
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)
	fmt.Println(s6)

}
