package nqueens

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func prepareMatrix(n int) io.Reader {
	var bb bytes.Buffer

	// Когда ферзь ставится на клетку, он занимает одну горизонталь, одну вертикаль и две диагонали (одна
	// слева направо и одна справа налево) на доске.
	// Соответственно, у нас будет всего 6n-2 итемов: n горизонталей, n вертикалей и по 2n-1 диагоналей.
	// Т.к. мы не можем требовать, чтобы каждая диагональ была занята ферзём, итемы для диагоналей делаем
	// вторичными.

	// Первичные итемы для вертикалей и горизонталей.
	items := make([]string, 0, 4*n-1)
	for i := 0; i <= n-1; i++ {
		items = append(items, fmt.Sprintf("x%d", i))
	}
	for i := 0; i <= n-1; i++ {
		items = append(items, fmt.Sprintf("y%d", i))
	}

	// Разделитель первичных и вторичных итемов.
	items = append(items, "|")

	// Вторичные итемы для диагоналей.
	for i := 0; i <= 2*n-2; i++ {
		items = append(items, fmt.Sprintf("zl%d", i))
	}
	for i := 0; i <= 2*n-2; i++ {
		items = append(items, fmt.Sprintf("zr%d", i))
	}

	bb.WriteString(strings.Join(items, " "))
	bb.WriteByte('\n')

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			bb.WriteString(fmt.Sprintf("x%d y%d zr%d zl%d\n", i, j, i+j, n-1-i+j))
		}
	}

	return &bb
}

func printBoard(s []string, n int) {
	board := make([][]byte, n)
	for i := range board {
		board[i] = make([]byte, n)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}

	var x, y int
	var err error

	for i := range s {
		opts := strings.Split(s[i], " ")
		for j := range opts {
			if opts[j][0] == 'x' {
				x, err = strconv.Atoi(opts[j][1:])
				if err != nil {
					panic(err)
				}
			}
			if opts[j][0] == 'y' {
				y, err = strconv.Atoi(opts[j][1:])
				if err != nil {
					panic(err)
				}
			}
		}
		board[x][y] = 'Q'
	}

	for i := range board {
		for j := range board[i] {
			fmt.Print(string(rune(board[i][j])), " ")
		}
		fmt.Println()
	}

	fmt.Println()
}
