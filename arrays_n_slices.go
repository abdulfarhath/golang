package main
import "fmt"

func main(){
	////////////declaring array

	// var array []int
	// arr := []int{}
	// arr := make([]int,5) //arr of size 5
	// arr[0] = 3

	//////////// filling array

	// arr := []int{} // puts it empty
	arr := []int{1,2,3} // fills with numbers 
	slice_types(arr);
}

func slice_types(arr []int){
	fmt.Println(arr)
}
