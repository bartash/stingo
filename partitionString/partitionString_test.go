package partitionString

import (
	"fmt"
	"testing"
)

func TestPartitionString(t *testing.T) {
	partitionString := NewPartitionString("and")
	expected := map[Pair]bool{
		*NewPair("a", "nd"): true,
		*NewPair("nd", "a"): true,
		*NewPair("n", "ad"): true,
		*NewPair("ad", "n"): true,
		*NewPair("an", "d"): true,
		*NewPair("d", "an"): true,
		*NewPair("and", ""): true,
		*NewPair("", "and"): true,
	}
	fmt.Printf("length of expect = %d\n", len(expected))
	for key, value := range expected {
		fmt.Print(key, "=", value)
	}
	pair := partitionString.Next()
	var count int
	for pair != nil {
		pair = partitionString.Next()
		if pair != nil {
			count++
			fmt.Printf("got pair '%v' & '%v'\n", pair.First, pair.Second)
			val, ok := expected[*pair]
			if !ok {
				t.Errorf("unexpected pair %+v %+v", pair, val)
			}
		} else {
			fmt.Printf("got nil return\n")
		}
	}
	if count != len(expected) {
		t.Errorf("unexpected count %d differs from len(expected)", count, len(expected))
	}

}
