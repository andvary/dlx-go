package nrooks

import (
	dlx "dlx_my"
	"testing"
)

func TestNRooks(t *testing.T) {
	m := prepareMatrix()
	d, err := dlx.New(m)
	if err != nil {
		t.Fatal(err)
	}

	if err := d.PrintSolutions(4); err != nil {
		t.Fatal(err)
	}
}
