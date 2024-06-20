package dlx

import (
	"os"
	"testing"
)

func TestDLX_Solve(t *testing.T) {
	f, err := os.OpenFile("./testdata/input_good_3.txt", os.O_RDONLY, 0111)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	d := &DLX{}

	if err := d.readInput(f); err != nil {
		t.Fatal(err)
	}

	if err := d.PrintSolutions(4); err != nil {
		t.Fatal(err)
	}
}
