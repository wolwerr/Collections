package main

import "fmt"

// func main() {

// 	arr:= [7]float64{1,2,3,4,5,6,7}
// 	fatia:= arr[0:5]
// 	    fmt.Println(fatia)
// }


// func main() {
//     fatia1:= []int{1,2,3}
// 	fatia2:= append(fatia1, 2, 4, 5)
// 	fmt.Println(fatia1, fatia2)
// }


func main() {
    fatia1:= []int{1,2,3}
	fatia2:= make([]int, 2)
	copy(fatia2, fatia1)
	fmt.Println(fatia1, fatia2)
}