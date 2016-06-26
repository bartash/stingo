package partitionString

import (
	"math"
)

type Pair struct {
	first, second string
}

type PartitionString struct {
	n    int
	seed string
}


func (p PartitionString) Next() *Pair {
	pair := new(Pair)
	max := int(math.Pow(2, float64(len(p.seed)))) // should be done in "constructor"
	if p.n > max {
		// need to check for off by one and other logic errors
		return nil
	}

	// treat n as a binary number to decide which characters of seed are in first string of pair

	p.n++
	return pair
}

