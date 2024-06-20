package nrooks

import (
	dlx "dlx_my"
	"testing"
)

func TestNRooks(t *testing.T) {
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
