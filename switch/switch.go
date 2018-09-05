package main

import "fmt"
import "time"

func main () {

	i := 2
	result := ""
	switch i {
	case 1:
		result = "one"
	case 2: 
		result = "two"
	case 3: 
		result = "three"
	}
	fmt.Println("Write", i, " as ", result)


	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")	
	}


	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a Boolean")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Println("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
