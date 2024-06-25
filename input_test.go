package dlx

import (
	"errors"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"os"
	"testing"
)

func TestReadInput(t *testing.T) {
	f, err := os.OpenFile("./testdata/input_good_3.txt", os.O_RDONLY, 0111)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	d := &DLX{}

	if err := d.readInput(f); err != nil {
		t.Fatal(err)
	}

	want := &DLX{
		primaryBoundary: 7,

		items: []*item{
			{name: "", prev: 7, next: 1},          //0
			{name: "a", prev: 0, next: 2, cnt: 2}, //1
			{name: "b", prev: 1, next: 3, cnt: 2}, //2
			{name: "c", prev: 2, next: 4, cnt: 2}, //3
			{name: "d", prev: 3, next: 5, cnt: 3}, //4
			{name: "e", prev: 4, next: 6, cnt: 2}, //5
			{name: "f", prev: 5, next: 7, cnt: 2}, //6
			{name: "g", prev: 6, next: 0, cnt: 3}, //7
		},
		opts: []*opt{
			{ //0
				prev:  6,
				next:  1,
				items: nil,
			},
			{ //1
				prev:  0,
				next:  2,
				items: []int{3, 5, 6},
			},
			{ //2
				prev:  1,
				next:  3,
				items: []int{1, 4, 7},
			},
			{ //3
				prev:  2,
				next:  4,
				items: []int{2, 3, 6},
			},
			{ //4
				prev:  3,
				next:  5,
				items: []int{1, 4},
			},
			{ //5
				prev:  4,
				next:  6,
				items: []int{2, 7},
			},
			{ //6
				prev:  5,
				next:  0,
				items: []int{4, 5, 7},
			},
		},
	}

	if diff := cmp.Diff(want, d, cmp.AllowUnexported(DLX{}, item{}, opt{}), cmpopts.IgnoreFields(opt{}, "lItems")); diff != "" {
		t.Errorf("want items mismatch(-want +got):\n%s", diff)
	}
}

func TestReadInput_WithOptionalItems(t *testing.T) {
	f, err := os.OpenFile("./testdata/input_good_3_optional_items.txt", os.O_RDONLY, 0111)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	d := &DLX{}

	if err := d.readInput(f); err != nil {
		t.Fatal(err)
	}

	want := &DLX{
		primaryBoundary: 5,
		items: []*item{
			{name: "", prev: 7, next: 1},          //0
			{name: "a", prev: 0, next: 2, cnt: 2}, //1
			{name: "b", prev: 1, next: 3, cnt: 2}, //2
			{name: "c", prev: 2, next: 4, cnt: 2}, //3
			{name: "d", prev: 3, next: 5, cnt: 3}, //4
			{name: "e", prev: 4, next: 6, cnt: 2}, //5
			{name: "f", prev: 5, next: 7, cnt: 2}, //6
			{name: "g", prev: 6, next: 0, cnt: 3}, //7
		},
		opts: []*opt{
			{ //0
				prev:  6,
				next:  1,
				items: nil,
			},
			{ //1
				prev:  0,
				next:  2,
				items: []int{3, 5, 6},
			},
			{ //2
				prev:  1,
				next:  3,
				items: []int{1, 4, 7},
			},
			{ //3
				prev:  2,
				next:  4,
				items: []int{2, 3, 6},
			},
			{ //4
				prev:  3,
				next:  5,
				items: []int{1, 4},
			},
			{ //5
				prev:  4,
				next:  6,
				items: []int{2, 7},
			},
			{ //6
				prev:  5,
				next:  0,
				items: []int{4, 5, 7},
			},
		},
	}

	if diff := cmp.Diff(want, d, cmp.AllowUnexported(DLX{}, item{}, opt{}), cmpopts.IgnoreFields(opt{}, "lItems")); diff != "" {
		t.Errorf("want items mismatch(-want +got):\n%s", diff)
	}
}

func TestReadInputBad(t *testing.T) {
	inputs := []string{
		"./testdata/input_bad_item_wt_opt.txt",
		"./testdata/input_bad_long_item_name.txt",
		"./testdata/input_bad_opt.txt",
		"./testdata/input_bad_0_primary_items.txt",
	}

	d := &DLX{}

	for i := range inputs {
		f, err := os.OpenFile(inputs[i], os.O_RDONLY, 0111)
		if err != nil {
			t.Fatal(err)
		}

		err = d.readInput(f)
		if err == nil {
			t.Errorf("функция не вернула ошибку для некорректных входных данных, файл %q", inputs[i])
		}

		if !errors.As(err, &InputError{}) {
			t.Errorf("ожидали ошибку типа InputError, получили: %v", err)
		}
		f.Close()
	}
}
