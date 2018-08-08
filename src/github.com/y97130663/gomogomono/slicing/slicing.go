package slicing

import "fmt"

var a int64 = 10
var b = []int64{2,3}

var slice1 = []int{3,4,5}
var slice2 = make([]int,2)

func Print_basic_slicing(){
	var x = [7]int64{1,2,3,4,5,6,7}
	var x_slicing_early = make([]int64, 5,7)
	x_slicing_later := x[6:len(x)]
	fmt.Println("x_slicing_early:", x_slicing_early)
	fmt.Println("x_slicing_later", x_slicing_later)
}

func Makes_sliceof10(a int64) []int64{
	x := make([]int64, a)
	return x
}

func Append_slicewith(b []int64) []int64{
	return append(b, 2, 3)
}

// Cannot append arrays since the size is fixed
//func Append_slicewithArray(d [5]int64) []int64{
//	return append(d, 2, 3)
//}

func Copy_Slicing(slice1 []int, slice2 []int){
	copy(slice2, slice1)
	fmt.Println(slice2)
}