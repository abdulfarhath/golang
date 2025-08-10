package main

import (
	"fmt"
)

func main(){
	var str string
	str = hi()
	fmt.Println(str)
}

func hi() string {
	return "hello this is my first go program"
}
