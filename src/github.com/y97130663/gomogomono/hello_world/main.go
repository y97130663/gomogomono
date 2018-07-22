package main

import "fmt"

type drive interface {
	carry() string
}

type vehicle struct{
	passengers, cargo bool
}

func (t vehicle) carry() string{
	if t.cargo && t.passengers {
		return "Yes it is truck"
	}
	return "No it is not a truck"
}

func main(){
	fmt.Println("Main function starts here!")
	t := vehicle {cargo: true, passengers: true}
	fmt.Println(t.carry())
}
