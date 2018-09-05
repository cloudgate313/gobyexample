package main

import "fmt"

func main () {

	// defining for loop variables
	i := 1
	for i <= 3 {
	fmt.Println(i)
	i = i + 1
	}
		
	// basic conditions
	for j := 7; j <= 9; j++ {
	fmt.Println(j)
	}	

	// break
	for {
	fmt.Println("Loop")
	fmt.Println("break jumps out of the for loop")
	break
	}

	// continue
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println("n%2 = 0 => continue jumps to next iteration")	
		fmt.Println(n)
	}
}
