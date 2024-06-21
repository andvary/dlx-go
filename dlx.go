package dlx

/*package dlx реализует алгоритм DLX Дональда Кнута.
Материалы: https://www-cs-faculty.stanford.edu/~knuth/programs.html
https://blog.demofox.org/2022/10/30/rapidly-solving-sudoku-n-queens-pentomino-placement-and-more-with-knuths-algorithm-x-and-dancing-links/
*/

import (
	"fmt"
	"io"
)

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

func New(r io.Reader, opts ...func(dlx *DLX)) (*DLX, error) {
	d := &DLX{}
	if err := d.readInput(r); err != nil {
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

func (d *DLX) SolveString() ([][]string, error) {
	if _, err := d.Solve(); err != nil {
		return nil, err
	}

	if len(d.solutions) == 0 {
		return nil, fmt.Errorf("no solutions found")
	}

	ss := make([][]string, len(d.solutions))

	for i, sol := range d.solutions {
		s := make([]string, len(sol))
		for j, op := range sol {
			s[j] = d.dumpOptions(op)
		}
		ss[i] = s
	}

	return ss, nil
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
