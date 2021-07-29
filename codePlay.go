package main

import(
	"fmt"
	//"strconv"
)

func sayGreeting(n string){
	fmt.Printf("Good morning %v \n", n)
}
func sayBye(n string){
	fmt.Printf("Good bye %v \n", n)
}

func cycleThrough(n []string, f func(string)){
	for _, i := range n {
		f(i)
	}
}

func main(){
	sayBye("kenneth")
	sayGreeting("nnah")
	cycleThrough([]string{"man", "woman", "gay"}, sayGreeting)
}