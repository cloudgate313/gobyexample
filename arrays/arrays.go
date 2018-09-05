package main

import "fmt"

func main() {

	// array with fixed length of 5
	var a [5]int
	fmt.Println("emp:", a)


	a[4] = 100
	fmt.Println("set: ", a)
	fmt.Println("get: ", a[4])

	fmt.Println("len: ", len(a))


	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl: ", b)


//      [amount of arrays][length of each array]
	var twoD [2][3]int
	for i := 0; i <  2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
//			fmt.Println("after ", twoD)
		}
	}
	fmt.Println("2d: ", twoD)
}
