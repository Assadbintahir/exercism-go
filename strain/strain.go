package strain

// Ints is a slice of integers
type Ints []int

// Lists is two dimensional slice of integers
type Lists [][]int

// Strings is a slice of strings
type Strings []string

func (intSlice Ints) Keep(pred func(int) bool) (result Ints) {
	for _, value := range intSlice {
		if pred(value) {
			result = append(result, value)
		}
	}
	return result
}

func (intSlice Ints) Discard(pred func(int) bool) (result Ints) {
	for _, value := range intSlice {
		if !pred(value) {
			result = append(result, value)
		}
	}
	return result
}

func (intLists Lists) Keep(pred func([]int) bool) (result Lists) {
	for _, value := range intLists {
		if pred(value) {
			result = append(result, value)
		}
	}
	return result
}

func (stringSlice Strings) Keep(pred func(string) bool) (result Strings) {
	for _, value := range stringSlice {
		if pred(value) {
			result = append(result, value)
		}
	}
	return result
}
