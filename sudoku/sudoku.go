package sudoku

import (
	"bytes"
	"dlx"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	cell    = "cel"
	row     = "row"
	column  = "col"
	square  = "sqr"
	initial = "i"
)

func solveSudoku(in [][]byte) {
	m := prepareMatrix(in)

	d, err := dlx.New(m)
	if err != nil {
		panic(err)
	}

	solutions, err := d.SolveString()
	if err != nil {
		panic(err)
	}

	for i := range solutions {
		for j := 1; j < len(solutions[i]); j++ {
			cellNum, err := strconv.Atoi(strings.TrimSpace(solutions[i][j][3:5]))
			if err != nil {
				panic(err)
			}
			ii, jj := (cellNum-1)/9, (cellNum-1)%9
			idx := strings.Index(solutions[i][j], "=")
			in[ii][jj] = solutions[i][j][idx+1]
		}
	}
}

func prepareMatrix(in [][]byte) io.Reader {
	bb := bytes.Buffer{}
	initialState := bytes.Buffer{}

	items := bytes.Buffer{}
	// 81 итем для значения (неважно какого) в каждой клетке
	for i := 1; i < 82; i++ {
		items.WriteString(fmt.Sprintf("%s%d ", cell, i))
	}
	// 81 итем для каждого возможно значения каждой из 9 колонок
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			items.WriteString(fmt.Sprintf("%s%d=%d ", column, i, j))
		}
	}
	// 81 итем для каждого возможно значения каждого из 9 рядов
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			items.WriteString(fmt.Sprintf("%s%d=%d ", row, i, j))
		}
	}
	// 81 итем для каждого возможно значения каждого из 9 квадратов
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			items.WriteString(fmt.Sprintf("%s%d=%d ", square, i, j))
		}
	}
	// Один итем для начального состояния (клетки, где цифры уже проставлены). Этот итем будет покрываться только одной
	// опцией, соответственно, всегда присутствовать в решении.
	items.WriteString(initial)

	items.WriteByte('\n')
	bb.Write(items.Bytes())

	for i := range in {
		for j := range in[i] {
			if in[i][j] == '.' {
				addUncoveredOption(&bb, i, j)
				continue
			}
			appendToInitialState(&initialState, string(in[i][j]), i, j)
		}
	}

	initialState.WriteString(initial)
	bb.Write(initialState.Bytes())

	return &bb
}

func addUncoveredOption(bb *bytes.Buffer, i, j int) {
	cellNum := 9*i + j + 1
	squareNum := 3*(i/3) + j/3 + 1
	colNum := j + 1
	rowNum := i + 1

	for i = 1; i < 10; i++ {
		bb.WriteString(fmt.Sprintf("%s%d %s%d=%d %s%d=%d %s%d=%d\n",
			cell, cellNum, column, colNum, i, row, rowNum, i, square, squareNum, i))
	}
}

func appendToInitialState(bb *bytes.Buffer, v string, i, j int) {
	cellNum := 9*i + j + 1
	squareNum := 3*(i/3) + j/3 + 1
	colNum := j + 1
	rowNum := i + 1

	bb.WriteString(fmt.Sprintf("%s%d %s%d=%s %s%d=%s %s%d=%s ",
		cell, cellNum, column, colNum, v, row, rowNum, v, square, squareNum, v))
}

func printBoard(in [][]byte) {
	for n := range in {
		fmt.Print("{")
		for nn := range in[n] {
			fmt.Printf("'%s',", string([]byte{in[n][nn]}))
		}
		fmt.Print("},\n")
	}
}
