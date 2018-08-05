package arrays

import "fmt"

var arraysExternal = [5]int64{
							3,
							4,
							5,
							6,
							7,
}
func Array_example_print() [5]int64{
	fmt.Println(arraysExternal)
	return arraysExternal
}
