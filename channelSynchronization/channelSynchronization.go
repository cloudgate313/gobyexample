package main

import "fmt"
import "time"

type input float64

func (i input) String() string {
	return fmt.Sprintf("%b", i)
}

func worker(done chan bool) {
	start := time.Now()

	fmt.Print("working...")
	fmt.Println("done")

	t := time.Now()
	fmt.Println(t.Sub(start))

	done <- true
}

func main() {
	
	done := make(chan bool, 1)
	go worker(done)

	<-done
}
