package partitionString

type Pair struct {
	first, second string
}

type PartitionString int

func NewPartitionString(s string) *PartitionString {
	ps  := new(PartitionString);
	return ps
}

func next() *Pair {
	return new(Pair)
}
