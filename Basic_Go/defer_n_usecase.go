package main
import "fmt"

func main(){
	defer fmt.Println("first")
	defer fmt.Println("second")
	fmt.Println("third")
}

//output
// third
// second 
// first

// defer executes just before function returns something
// till then the function with defer waits for its time to come
// like me waiting to get a job

// it follows a lifo technique btw
// the last input defer function executes first
// like the lucky people get job first 

