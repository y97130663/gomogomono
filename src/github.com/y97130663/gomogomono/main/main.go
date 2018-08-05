package main

import (
	"fmt"
	"github.com/y97130663/gomogomono/primenumber"
	"reflect"
	"github.com/y97130663/gomogomono/slicing"
	"github.com/y97130663/gomogomono/arrays"
)

var k = 0
var input string

func main(){
	var a = k
	fmt.Println("The variable a is a package level scope")
	fmt.Printf("%v\n", a)
	primenumber.Primenumbers(a)
	fmt.Scan( &input)
	//primenumber.Primenumbers(B)
	fmt.Printf("The user entered value with improper string formatting is %s \n", input)
	fmt.Printf("The user entered value is %d \n", input)
	fmt.Printf("The user entered value after type conversion is %d \n", input)
	x := 2
	var y [5]int
	for i := 0; i <= x; i++{
		y[i] = i
	}
	fmt.Println(y)
	fmt.Println(y[2])

	//performing fizzbuzz
	fizzbuzz()

	//slicing
	slicing.Print_basic_slicing()
	c := slicing.Makes_sliceof10(10)
	d := slicing.Append_slicewith(c)
	fmt.Println("The 10x1 array is:", c)
	fmt.Println("The new slice appended is:", d)

	//Setter arrays
	fmt.Println("h is:", arrays.Array_example_print_setter())

	//Another way to print
	arrays.Print_ArrayExternal()

}


func fizzbuzz(){
	/*
	Write a program that prints the numbers from 1 to 100.
	But for multiples of three print "Fizz" in- stead
	of the number and for the multiples of five print "Buzz".
	For numbers which are multiples of both three and five print "FizzBuzz".
	*/
	input2 := 10
	fmt.Println("Enter the range for the fizzbuzz")
	fmt.Scan( &input2)
	fmt.Printf("Enter the range for the fizzbuzz is %d \n", input2)
	fmt.Println("The format is ", reflect.TypeOf(input2))
	for i:=0; i<=int(input2); i++{
		if i % 3 == 0 && i % 15 !=0{
			fmt.Printf("%d Fizz \n", i)
		} else if i % 5 ==0 && i % 15 != 0{
			fmt.Printf("%d Buzz \n", i)
		} else if i % 15 == 0{
			fmt.Printf("%d FizzBuzz \n", i)
		}
	}

}