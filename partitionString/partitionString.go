package partitionString

type Pair struct {
	first, second string
}

type PartitionString struct {
	N int
	S string
}


func (p PartitionString) Next() *Pair {
	return new(Pair)
}

