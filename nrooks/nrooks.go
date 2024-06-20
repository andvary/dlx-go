package nrooks

import (
	"bytes"
	"io"
	"strings"
)

func prepareMatrix() io.Reader {
	var bb bytes.Buffer

	items := []string{"x1", "x2", "x3", "x4", "x5", "x6", "x7", "x8", "y1", "y2", "y3", "y4", "y5", "y6", "y7", "y8"}

	bb.WriteString(strings.Join(items, " "))
	bb.WriteByte('\n')

	for i := 0; i < 8; i++ {
		for j := 8; j < len(items); j++ {
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
