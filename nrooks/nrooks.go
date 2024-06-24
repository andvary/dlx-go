package nrooks

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func prepareMatrix(x, y int) io.Reader {
	var bb bytes.Buffer

	// Когда ладья ставится на клетку, она занимает одну горизонталь и одну вертикаль на доске.
	// Соответственно, у нас будет всего 16 итемов: 8 для горизонталей и 8 для вертикалей.
	items := make([]string, 0, x+y)
	for i := 1; i <= x; i++ {
		items = append(items, fmt.Sprintf("x%d", i))
	}
	for i := 1; i <= y; i++ {
		items = append(items, fmt.Sprintf("y%d", i))
	}

	bb.WriteString(strings.Join(items, " "))
	bb.WriteByte('\n')

	// Всего 64 опции, для каждой клетки, на которую можно поставить ладью. Если ладья ставится на клетку (0,0),
	// добавляем опцию "X0 Y0". Если на (4,5) - "X4 Y5"
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

func printBoard(x, y int, solutions [][]int, r io.Reader) {
	board := make([][]byte, y)

	for i := range board {
		row := make([]byte, x)
		for j := range row {
			row[j] = '.'
		}
		board[i] = row
	}

	options := make([]string, 0)
	s := bufio.NewScanner(r)
	for s.Scan() {
		options = append(options, s.Text())
	}

	var xx, yy int
	var err error

	for _, solution := range solutions {
		for _, opt := range solution {
			ss := strings.Split(options[opt], " ")
			for j := range ss {
				if ss[j][0] == 'x' {
					xx, err = strconv.Atoi(ss[j][1:])
					if err != nil {
						panic(err)
					}
				}
				if ss[j][0] == 'y' {
					yy, err = strconv.Atoi(ss[j][1:])
					if err != nil {
						panic(err)
					}
				}
			}
			board[xx-1][yy-1] = 'X'
		}

		for i := range board {
			for j := range board[i] {
				fmt.Printf("%s ", string(board[i][j]))
			}
			fmt.Println()
		}
		fmt.Println()

		// очистим поле
		for i := range board {
			row := make([]byte, x)
			for j := range row {
				row[j] = '.'
			}
			board[i] = row
		}
	}
}
