package mydlx

import "fmt"

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

	debug bool
}

func NewFromFile(path string, opts ...func(dlx *DLX)) (*DLX, error) {
	d := &DLX{}
	if err := d.readInput(path); err != nil {
		return nil, err
	}

	for i := range opts {
		opts[i](d)
	}

	return d, nil
}

func EnableDebugging() func(dlx *DLX) {
	return func(d *DLX) {
		d.debug = true
	}
}

func (d *DLX) Solve() ([][]int, error) {
	bestItem := d.findBestItem()
	if err := d.cover(bestItem); err != nil {
		return nil, err
	}
	return d.solutions, nil
}

func (d *DLX) PrintSolutions(maxSolutions int) error {
	if _, err := d.Solve(); err != nil {
		return err
	}

	if len(d.solutions) == 0 {
		fmt.Println("no solutions found")
		return nil
	}

	fmt.Printf("\n\n%d solution(s) total\n\n", len(d.solutions))

	for i := 0; i < maxSolutions && i < len(d.solutions); i++ {
		fmt.Printf("solution %d:\n", i+1)
		fmt.Println(d.dumpOptions(d.solutions[i]...))
	}

	return nil
}
