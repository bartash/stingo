package partitionString

import (
	"testing"
	"fmt"
)

func TestPartitionString(t *testing.T) {
	out := PartitionString{0, "and"}

	pair := out.Next();
	fmt.Printf("got pair '%v' & '%v'\n", pair.first, pair.second)

	for pair != nil {
		pair = out.Next();
		if pair != nil {
			fmt.Printf("got pair '%v' & '%v'\n", pair.first, pair.second)
		} else {
			fmt.Printf("got nil return\n")
		}
	}
}
