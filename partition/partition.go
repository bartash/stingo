package partition


type Pair struct {
	first, second string
}


// return all strings pairs formed from all the characters in the string
// where the length of the first string is as specified
func Partition(s string, length int) []Pair {

	if length == 1 {
		var  ret []Pair
		stringLength := len(s)

		for i := 0; i < stringLength; i++ {
			var left string = ""
			var right string = ""
			for j := 0; j < stringLength; j++ {
				if i == j {
					left = left + string(s[j])
				} else {
					right = right + string(s[j])
				}
			}
			p := Pair{left, right}
			ret = append(ret, p)
		}
		return ret
	}
	panic("not implemented")
}