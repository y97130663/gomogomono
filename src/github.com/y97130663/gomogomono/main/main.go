package main

import (
	"fmt"
	"github.com/y97130663/gomogomono/primenumber"
)

var k = 0
var input int16

func main(){
	var a = k
	fmt.Println("The variable a is a package level scope")
	fmt.Printf("%v\n", a)
	primenumber.Primenumbers(a)
	fmt.Scan( &input)
	//primenumber.Primenumbers(B)
	fmt.Printf("The user entered value with improper string formatting is %s \n", input)
	fmt.Printf("The user entered value is %d \n", input)
	fmt.Printf("The user entered value after type conversion is %d \n", int(input))

	x := 2
	var y [5]int

	for i := 0; i <= x; i++{
		y[i] = i
	}

	fmt.Println(y)
	fmt.Println(y[2])

	fizzbuzz()
}

func fizzbuzz(){
	/*
	Writeaprogramthatprintsthenumbersfrom1 to 100.
	But for multiples of three print "Fizz" in- stead
	of the number and for the multiples of five print "Buzz".
	For numbers which are multiples of both three and five print "FizzBuzz".
	*/

	var input2 int32
	fmt.Println("Enter the range for the fizzbuzz")
	fmt.Scan( &input2)
	fmt.Printf("Enter the range for the fizzbuzz is %d", input2)


}