package partitionString

import (
	"testing"
	"fmt"
)

func TestPartitionString(t *testing.T) {
	out := PartitionString{0, "and"}
	pair := out.Next();
	fmt.Println(pair)
}
