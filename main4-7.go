package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		
		fmt.Println(i)

		if i == 0 {
			
		} else if i == 1{
			
		}else{

		}

		switch i {
			case 0: fmt.Println("Zero")
			case 1: fmt.Println("One")
			case 2: fmt.Println("Two")
			case 3: fmt.Println("Three")
			case 4: fmt.Println("Four")
			case 5: fmt.Println("Five")
			default: fmt.Println("Unknown Number")
		}
	}

	basicTypes2()
}

func basicTypes2(){

	// Arrays
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	var total float64 = 0

	for i := 0; i < len(x); i++ {
		total += x[i]
	}

	x2 := [4]float64{ 98, 93, 77,
					  // 82,
					  83 }

	for _, value := range x2 {
		total += value
	}

	fmt.Println(total / float64(len(x)))

	/* Slices -

		- They are always associated with some array 
		- They can never be longer than the array, they can be smaller.
		
	*/

}