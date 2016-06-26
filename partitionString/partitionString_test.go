package partitionString

import (
	"testing"
	"fmt"
)

func TestPartitionString(t *testing.T) {
	second := NewPartitionString("and")
	pair := second.Next()
	for pair != nil {
		pair = second.Next();
		if pair != nil {
			fmt.Printf("got pair '%v' & '%v'\n", pair.First, pair.Second)
		} else {
			fmt.Printf("got nil return\n")
		}
	}
}
