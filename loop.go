package main

import(
	"fmt"
)


func main(){
//  var name string = "Nnamdi"
// 	for i := 0; i < len(name); i++ {
// 		fmt.Println(i)
// 	}
 //fmt.Println(len(name)) 

 stringSlice := []string{"man", "woman", "frodite"}
 for i := 0; i < len(stringSlice); i++ {            // Looping through a slice of array
	 fmt.Println((stringSlice[i]))
 }

 countries := []string{"Nigeria", "mali", "Bulgaria"} // Slice of array in go
 	for index, value := range countries {
		 fmt.Printf("output is %v %v\n", index, value)
	 }
}