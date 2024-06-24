package nrooks

import (
	"dlx"
	"testing"
)

func TestNRooks_8(t *testing.T) {
	m := prepareMatrix(8, 8)
	d, err := dlx.New(m)
	if err != nil {
		t.Fatal(err)
	}

	solutions, err := d.Solve()
	if err != nil {
		t.Fatal(err)
	}

	if len(solutions) != 40320 {
		t.Fatal(err)
	}

	printBoard(8, 8, solutions[3100:3104], prepareMatrix(8, 8))
}

func TestNRooks_10(t *testing.T) {
	m := prepareMatrix(10, 10)
	d, err := dlx.New(m)
	if err != nil {
		t.Fatal(err)
	}

	solutions, err := d.Solve()
	if err != nil {
		t.Fatal(err)
	}

	if len(solutions) != 3628800 {
		t.Fatal(err)
	}

	printBoard(10, 10, solutions[3100:3104], prepareMatrix(10, 10))
}
