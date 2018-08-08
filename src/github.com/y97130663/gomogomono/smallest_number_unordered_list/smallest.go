package smallest_number_unordered_list

var x = []int{48,96,86,68,57,82,63,70,37,34,83,27,19,97,9,17,}
const j = 0

func Smallest() int{
	for i := 0; i < len(x)-1; i++  {
		if x[j] < x [i+1]{
			continue
		} else{
			x[j] = x[i+1]
		}
	}
	return x[j]
}

func Smallest_with_range() int{
	for _, ranged := range x{
		if x[j] < ranged{
			continue
		} else{
			x[j] = ranged
		}
	}
	return x[j]
}
