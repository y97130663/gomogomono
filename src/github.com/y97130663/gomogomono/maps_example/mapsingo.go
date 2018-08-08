package maps_example

import "fmt"

var Y map[string]int

func Print_maps(){
	X := make(map[string]int)
	X["key"] = 233
	fmt.Println(X["key"])

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"
	fmt.Println(elements["Li"])
	fmt.Println(elements)
	new_Map()
	nested_Maps()

	if k, ok := elements["N"]; ok {
		fmt.Println("The element exists and it is ", k)
	}

	if g, ok := elements["Lit"]; ok {
		fmt.Println("The second element exists", g)
	} else {
		fmt.Println("The element does not exist")
	}

}

func new_Map(){
	new_map := map[string]int{
		"a" : 1,
		"b" : 2,
		"c" : 3,
	}
	fmt.Println(new_map)
}

func nested_Maps(){
	nested_new_map := map[string]map[string]int{
		"A" : map[string]int{
			"a": 1,
		},
		"B" : {
			"b": 2,
		},
	}
	fmt.Println(nested_new_map)

}