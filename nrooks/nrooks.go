package nrooks

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func prepareMatrix(x, y int) io.Reader {
	var bb bytes.Buffer

	items := make([]string, 0, x+y)
	for i := 1; i <= x; i++ {
		items = append(items, fmt.Sprintf("x%d", i))
	}
	for i := 1; i <= y; i++ {
		items = append(items, fmt.Sprintf("y%d", i))
	}

	bb.WriteString(strings.Join(items, " "))
	bb.WriteByte('\n')

	for i := 0; i < x; i++ {
		for j := x; j < len(items); j++ {
			bb.WriteString(items[i])
			bb.WriteByte(' ')
			bb.WriteString(items[j])
			bb.WriteByte('\n')
		}
	}

	return &bb
}

func printBoard(s []int, reader io.Reader) {

}
