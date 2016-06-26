package partitionString

import (
	"math"
)

type Pair struct {
	first, second string
}

type PartitionString struct {
	n    int  // FIXME should 64 bits and this limits the strings we can use
	seed string
}


func (p PartitionString) Next() *Pair {
	pair := new(Pair)
	stringLength := len(p.seed)
	max := int(math.Pow(2, float64(stringLength))) // should be done in "constructor"
	if p.n > max {
		// need to check for off by one and other logic errors
		return nil
	}

	// treat n as a binary number to decide which characters of seed are in first string of pair
	for i := 0; i < stringLength; i++ {
		c := p.seed[i]
		if isBitSet(p.n, i) {
			pair.first = append(pair.first, c)
		} else {
			pair.second = append(pair.second, c)
		}

	}

	p.n++
	return pair
}

func isBitSet(n int, pos uint) bool {
	val := n & (1 << pos)
	return (val > 0)
}

