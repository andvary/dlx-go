package nrooks

import (
	dlx "dlx_my"
	"testing"
)

func TestNRooks(t *testing.T) {
	m := prepareMatrix(2, 2)
	d, err := dlx.New(m, dlx.EnableDebugging())
	if err != nil {
		t.Fatal(err)
	}

	if err := d.PrintSolutions(4); err != nil {
		t.Fatal(err)
	}
}
