package mydlx

import "testing"

func TestDLX_Solve(t *testing.T) {
	d, err := NewFromFile("./testdata/input_good_3.txt")
	if err != nil {
		t.Fatal(err)
	}

	if err := d.PrintSolutions(4); err != nil {
		t.Fatal(err)
	}
}
