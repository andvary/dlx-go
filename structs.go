package mydlx

type opt struct {
	prev  int
	next  int
	items map[int]struct{}
}

type item struct {
	name string
	prev int
	next int
	cnt  int
}

type DLX struct {
	items []*item
	opts  []*opt

	solutions         [][]int
	potentialSolution []int
}

func NewFromFile(path string) *DLX {
	return nil
}
