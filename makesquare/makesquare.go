package makesquare

import (
	"bytes"
	"dlx"
	"io"
	"sort"
	"strconv"
)

// makesquare дано: набор спичек, длины спичек представлены элементами в массиве m.
// Нужно установить, можно ли, выложить из предложенных спичек квадрат. Ломать спички нельзя, соединять можно.
func makesquare(m []int) bool {
	if len(m) < 4 {
		return false
	}

	var sum int
	for i := range m {
		sum += m[i]
	}

	if sum%4 != 0 {
		return false
	}

	for i := range m {
		if m[i] > sum/4 {
			return false
		}
	}

	matrix := prepareMatrix(m, sum/4)

	d, err := dlx.New(matrix, dlx.MaxSolutions(1))
	if err != nil {
		panic(err)
	}

	res, err := d.Solve()
	if err != nil {
		panic(err)
	}

	return len(res) > 0
}

func prepareMatrix(m []int, side int) io.Reader {
	bb := bytes.Buffer{}

	// Начинаем с самых длинных опций, неимоверно ускоряет работу.
	sort.Ints(m)

	for i := range m {
		bb.WriteByte('m')
		bb.WriteString(strconv.Itoa(i))
		bb.WriteByte(' ')
	}

	for i := 0; i < side*4; i++ {
		bb.WriteByte('p')
		bb.WriteString(strconv.Itoa(i))
		bb.WriteByte(' ')
	}

	bb.WriteByte('\n')

	sides := [4]int{0, side, side * 2, side * 3}

	for i := range m {

		positions := getPositions(m[i], side)
		for _, s := range sides {
			for _, pp := range positions {
				bb.WriteByte('m')
				bb.WriteString(strconv.Itoa(i))
				bb.WriteByte(' ')

				for _, p := range pp {
					bb.WriteByte('p')
					bb.WriteString(strconv.Itoa(p + s))
					bb.WriteByte(' ')
				}

				bb.WriteByte('\n')
			}
		}

	}

	return &bb
}

func getPositions(m int, side int) [][]int {
	res := make([][]int, 0, side-m+1)

	r := make([]int, m)
	for i := range r {
		r[i] = i
	}

	for r[len(r)-1] < side {
		res = append(res, make([]int, len(r)))
		copy(res[len(res)-1], r)
		for i := range r {
			r[i]++
		}
	}

	return res
}
