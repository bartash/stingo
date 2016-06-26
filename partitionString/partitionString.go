package partitionString

import (
	"math"
)

type Pair struct {
	First, Second string
}

func NewPair(first string, second string) * Pair {
	p := new(Pair)
	p.First = first
	p.Second = second
	return p
}

type PartitionString struct {
	n    int  // FIXME should 64 bits and this limits the strings we can use
	seed string
	max int
}

func NewPartitionString(s string) * PartitionString {
	p := new(PartitionString)
	p.n = 0;
	p.seed = s
	p.max = int(math.Pow(2, float64(len(s))))
	return p;
}


func (p * PartitionString) Next() *Pair {
	pair := new(Pair)
	stringLength := len(p.seed)

	// fmt.Printf("max=%v n=%v\n", max, p.n)
	if p.n >= p.max {
		// need to check for off by one and other logic errors
		return nil
	}

	// treat n as a binary number to decide which characters of seed are in first string of pair
	for i := 0; i < stringLength; i++ {
		c := string(p.seed[i])
		// fmt.Printf("i=%v n=%v\n", i, p.n)
		if isBitSet(p.n, uint(i)) {
			pair.First += c
		} else {
			pair.Second += c
		}
	}

	p.n++
	// fmt.Printf("inc n=%v\n", p.n)
	return pair
}

func isBitSet(n int, pos uint) bool {
	val := n & (1 << pos)
	return (val > 0)
}

