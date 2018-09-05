package main

import "fmt"

func main() {

	var a = "verbose variable declaration"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b,c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "concise variable declaration"
	fmt.Println(f)

}
