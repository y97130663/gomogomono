package arrays

import "fmt"

var arraysExternal = [5]int64{
	3,
	4,
	5,
	6,
	7,
}
func Array_example_print_setter() [5]int64{
	return arraysExternal
}

func Print_ArrayExternal(){
	fmt.Println("arraysExternal:", arraysExternal)
}

