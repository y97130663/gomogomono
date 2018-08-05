package slicing

import "fmt"

func Print_basic_slicing(){
	var x = [7]int64{1,2,3,4,5,6,7}
	var x_slicing_early = make([]int64, 5,7)
	x_slicing_later := x[6:len(x)]
	fmt.Println("x_slicing_early:", x_slicing_early)
	fmt.Println("x_slicing_later", x_slicing_later)
}
