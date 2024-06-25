package makesquare

import (
	"bytes"
	"dlx"
	"io"
	"log"
	"strconv"
	"strings"
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

	// Сначала найдём все возможные варианты получения отрезка длиной side из имеющихся спичек.
	matrix := prepareMatrix(m, sum/4)

	d, err := dlx.New(matrix)
	if err != nil {
		log.Print(err)
		return false
	}

	subsets, err := d.SolveString()
	if err != nil {
		log.Print(err)
		return false
	}

	// Потом выясним, можно ли, используя полученные выше варианты, сложить квадрат, используя при этом все спички.
	d, err = dlx.New(prepareMatrix2(m, subsets), dlx.MaxSolutions(1))
	if err != nil {
		log.Print(err)
		return false
	}

	res, err := d.Solve()
	if err != nil {
		log.Print(err)
		return false
	}

	return len(res) > 0
}

func prepareMatrix(m []int, side int) io.Reader {
	bb := bytes.Buffer{}

	for i := 0; i < side; i++ {
		bb.WriteByte('p')
		bb.WriteString(strconv.Itoa(i))
		bb.WriteByte(' ')
	}

	bb.WriteByte('|')
	bb.WriteByte(' ')

	for i := range m {
		bb.WriteByte('m')
		bb.WriteString(strconv.Itoa(i))
		bb.WriteByte(' ')
	}

	bb.WriteByte('\n')

	for i := range m {

		positions := getPositions(m[i], side)

		for _, pp := range positions {
			bb.WriteByte('m')
			bb.WriteString(strconv.Itoa(i))
			bb.WriteByte(' ')

			for _, p := range pp {
				bb.WriteByte('p')
				bb.WriteString(strconv.Itoa(p))
				bb.WriteByte(' ')
			}

			bb.WriteByte('\n')
		}
	}

	return &bb
}

func prepareMatrix2(m []int, subsets [][]string) io.Reader {
	bb := bytes.Buffer{}

	for i := range m {
		bb.WriteByte('m')
		bb.WriteString(strconv.Itoa(i))
		bb.WriteByte(' ')
	}
	bb.WriteByte('\n')

	var idx int
	for i := range subsets {
		for j := range subsets[i] {
			idx = strings.Index(subsets[i][j], " ")
			bb.WriteString(subsets[i][j][:idx])
			bb.WriteByte(' ')
		}
		bb.WriteByte('\n')
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
