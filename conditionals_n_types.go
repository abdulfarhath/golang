package main

import "fmt"

func main(){
	//if condition

	age := 12

	if(age > 18){
		fmt.Println("adult")
	}else if(age == 18){
		fmt.Println("just 18")
	}else{
		fmt.Printnl("child")
	}

	switch(age){
		case 18:
			fmt.Println("just 18")
			break;
		case 18...20:
			fmt.Println("adult")
			break;
		default:
			fmt.Println("enter correct age")
	}
}
