package main
import (
	"fmt"
	"os"
	"io"
) 

func main(){
	str := "i use arch btw"

	//create a file
	file,err := os.Create("./manipulate_me.txt")
	if(err != nil){
		panic(err) // panic if error
	}

	//writes str into file
	length, err := io.WriteString(file,str)
	panicForErr(err)

	fmt.Println("length of string : ",length);
	readFile(file)

	//close file
	defer file.Close() //defines how responsible i am,😂 
}

func readFile(file string){
	
	//reading a file
	str,err := ioutil.Readfile(file)
	panicForErr(err)

	fmt.Println(string(str))
}

//wrote a panic method instead
func panicForErr(err Error){
	if(err != nil){
		panic(err)
	}
}
