package main

import(
	"fmt"
	"strconv"
)


func main(){
	// convert the below string to have quotes
	myString := "The delegence of a man lifts him up"
	fmt.Println(strconv.Quote(myString))
}