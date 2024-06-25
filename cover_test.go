package dlx

import (
	"errors"
	"github.com/google/go-cmp/cmp"
	"os"
	"sort"
	"strconv"
	"testing"
)

func TestRemoveItem_Good(t *testing.T) {
	tt := map[string]struct {
		in           []*item
		itemToRemove int
		want         []*item
	}{
		"1": {
			in: []*item{
				{name: "", prev: 4, next: 1},          //0
				{name: "a", prev: 0, next: 2, cnt: 1}, //1
				{name: "b", prev: 1, next: 3, cnt: 2}, //2
				{name: "c", prev: 2, next: 4, cnt: 2}, //3
				{name: "d", prev: 3, next: 0, cnt: 1}, //4
			},
			itemToRemove: 1,
			want: []*item{
				{name: "", prev: 4, next: 2},          //0
				{name: "a", prev: 0, next: 2, cnt: 1}, //1
				{name: "b", prev: 0, next: 3, cnt: 2}, //2
				{name: "c", prev: 2, next: 4, cnt: 2}, //3
				{name: "d", prev: 3, next: 0, cnt: 1}, //4
			},
		},
		"2": {
			in: []*item{
				{name: "", prev: 4, next: 1},          //0
				{name: "a", prev: 0, next: 2, cnt: 1}, //1
				{name: "b", prev: 1, next: 3, cnt: 2}, //2
				{name: "c", prev: 2, next: 4, cnt: 2}, //3
				{name: "d", prev: 3, next: 0, cnt: 1}, //4
			},
			itemToRemove: 3,
			want: []*item{
				{name: "", prev: 4, next: 1},          //0
				{name: "a", prev: 0, next: 2, cnt: 1}, //1
				{name: "b", prev: 1, next: 4, cnt: 2}, //2
				{name: "c", prev: 2, next: 4, cnt: 2}, //3
				{name: "d", prev: 2, next: 0, cnt: 1}, //4
			},
		},
		"3": {
			in: []*item{
				{name: "", prev: 4, next: 1},          //0
				{name: "a", prev: 0, next: 2, cnt: 1}, //1
				{name: "b", prev: 1, next: 3, cnt: 2}, //2
				{name: "c", prev: 2, next: 4, cnt: 2}, //3
				{name: "d", prev: 3, next: 0, cnt: 1}, //4
			},
			itemToRemove: 4,
			want: []*item{
				{name: "", prev: 3, next: 1},          //0
				{name: "a", prev: 0, next: 2, cnt: 1}, //1
				{name: "b", prev: 1, next: 3, cnt: 2}, //2
				{name: "c", prev: 2, next: 0, cnt: 2}, //3
				{name: "d", prev: 3, next: 0, cnt: 1}, //4
			},
		},
		"4": {
			in: []*item{
				{name: "", prev: 1, next: 1},          //0
				{name: "a", prev: 0, next: 0, cnt: 1}, //1
			},
			itemToRemove: 1,
			want: []*item{
				{name: "", prev: 0, next: 0},          //0
				{name: "a", prev: 0, next: 0, cnt: 1}, //1
			},
		},
	}

	for name, tc := range tt {
		d := &DLX{items: make([]*item, len(tc.in))}
		copy(d.items, tc.in)
		t.Run(name, func(t *testing.T) {
			if err := d.removeItem(tc.itemToRemove); err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(tc.want, d.items, cmp.AllowUnexported(item{})); diff != "" {
				t.Errorf("want items mismatch(-want +got):\n%s", diff)
			}
		})
	}
}

func TestRemoveItem_Bad(t *testing.T) {
	in := []*item{
		{name: "", prev: 4, next: 1},          //0
		{name: "a", prev: 0, next: 2, cnt: 1}, //1
		{name: "b", prev: 1, next: 3, cnt: 2}, //2
		{name: "c", prev: 2, next: 4, cnt: 2}, //3
		{name: "d", prev: 3, next: 0, cnt: 1}, //4
	}

	indices := []int{0, len(in), len(in) + 1, -1}

	for _, index := range indices {
		d := &DLX{items: make([]*item, len(in))}
		copy(d.items, in)
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			err := d.removeItem(index)
			if err == nil {
				t.Errorf("функция не вернула ошибку при удалении итема %d", index)
			}

			if !errors.As(err, &CoverError{}) {
				t.Fatalf("ожидали CoverError получили %v", err)
			}
		})
	}
}

func TestRemoveOption(t *testing.T) {
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
			{name: "c", prev: 2, next: 4, cnt: 1}, //3
			{name: "d", prev: 3, next: 5, cnt: 3}, //4
			{name: "e", prev: 4, next: 6, cnt: 1}, //5
			{name: "f", prev: 5, next: 7, cnt: 1}, //6
			{name: "g", prev: 6, next: 0, cnt: 3}, //7
		},
		opts: []*opt{
			{
				prev:  6,
				next:  2,
				items: nil,
			},
			{
				prev:  0,
				next:  2,
				items: []int{3, 5, 6},
			},
			{
				prev:  0,
				next:  3,
				items: []int{1, 4, 7},
			},
			{
				prev:  2,
				next:  4,
				items: []int{2, 3, 6},
			},
			{
				prev:  3,
				next:  5,
				items: []int{1, 4},
			},
			{
				prev:  4,
				next:  6,
				items: []int{2, 7},
			},
			{
				prev:  5,
				next:  0,
				items: []int{4, 5, 7},
			},
		},
	}

	if err := d.removeOption(1); err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(want, d, cmp.AllowUnexported(DLX{}, item{}, opt{})); diff != "" {
		t.Errorf("want items mismatch(-want +got):\n%s", diff)
	}
}

func TestCover(t *testing.T) {
	f, err := os.OpenFile("./testdata/input_good_3.txt", os.O_RDONLY, 0111)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	d := &DLX{}

	if err := d.readInput(f); err != nil {
		t.Fatal(err)
	}

	if err := d.cover(1); err != nil {
		t.Fatal(err)
	}

	want := [][]int{
		{1, 4, 5},
	}

	for i := range d.solutions {
		sort.Ints(d.solutions[i])
	}

	if diff := cmp.Diff(want, d.solutions); diff != "" {
		t.Errorf("want items mismatch(-want +got):\n%s", diff)
	}
}

func TestFindBestItem(t *testing.T) {
	d := &DLX{
		items: []*item{
			{name: "", prev: 6, next: 2},          //0
			{name: "a", prev: 0, next: 2, cnt: 1}, //1 x
			{name: "b", prev: 0, next: 3, cnt: 4}, //2
			{name: "c", prev: 2, next: 5, cnt: 4}, //3
			{name: "d", prev: 3, next: 5, cnt: 2}, //4 x
			{name: "e", prev: 3, next: 6, cnt: 5}, //5
			{name: "f", prev: 5, next: 0, cnt: 5}, //6
			{name: "g", prev: 6, next: 0, cnt: 3}, //7 x
		},
	}

	got := d.findBestItem()
	want := 2

	if got != want {
		t.Fatalf("ожидали: %d, получили: %d", want, got)
	}
}

func TestRestoreItems(t *testing.T) {
	d := &DLX{
		items: []*item{
			{name: "", prev: 0, next: 0},          //0
			{name: "a", prev: 0, next: 2, cnt: 2}, //1	x
			{name: "b", prev: 0, next: 3, cnt: 2}, //2	x
			{name: "c", prev: 0, next: 0, cnt: 2}, //3	x
			{name: "d", prev: 3, next: 0, cnt: 3}, //4	x
			{name: "e", prev: 4, next: 0, cnt: 2}, //5	x
			{name: "f", prev: 5, next: 0, cnt: 2}, //6	x
			{name: "g", prev: 6, next: 0, cnt: 3}, //7	x
		},
	}

	want := []*item{
		{name: "", prev: 7, next: 1},          //0
		{name: "a", prev: 0, next: 2, cnt: 2}, //1
		{name: "b", prev: 1, next: 3, cnt: 2}, //2
		{name: "c", prev: 2, next: 4, cnt: 2}, //3
		{name: "d", prev: 3, next: 5, cnt: 3}, //4
		{name: "e", prev: 4, next: 6, cnt: 2}, //5
		{name: "f", prev: 5, next: 7, cnt: 2}, //6
		{name: "g", prev: 6, next: 0, cnt: 3}, //7
	}

	removedItems := []int{1, 7, 6, 5, 2, 4, 3}
	d.restoreItems(removedItems)

	if diff := cmp.Diff(want, d.items, cmp.AllowUnexported(item{})); diff != "" {
		t.Errorf("want items mismatch(-want +got):\n%s", diff)
	}

}

func TestRestoreItems2(t *testing.T) {
	d := &DLX{
		items: []*item{
			{name: "", prev: 6, next: 4},          //0
			{name: "a", prev: 0, next: 2, cnt: 2}, //1	xx
			{name: "b", prev: 0, next: 3, cnt: 2}, //2	x
			{name: "c", prev: 0, next: 4, cnt: 2}, //3	x
			{name: "d", prev: 0, next: 6, cnt: 3}, //4
			{name: "e", prev: 4, next: 6, cnt: 2}, //5	xx
			{name: "f", prev: 4, next: 0, cnt: 2}, //6
			{name: "g", prev: 6, next: 0, cnt: 3}, //7	x
		},
	}

	want := []*item{
		{name: "", prev: 7, next: 2},          //0
		{name: "a", prev: 0, next: 2, cnt: 2}, //1	xx
		{name: "b", prev: 0, next: 3, cnt: 2}, //2
		{name: "c", prev: 2, next: 4, cnt: 2}, //3
		{name: "d", prev: 3, next: 6, cnt: 3}, //4
		{name: "e", prev: 4, next: 6, cnt: 2}, //5	xx
		{name: "f", prev: 4, next: 7, cnt: 2}, //6
		{name: "g", prev: 6, next: 0, cnt: 3}, //7
	}

	removedItems := []int{2, 3, 7}
	d.restoreItems(removedItems)

	if diff := cmp.Diff(want, d.items, cmp.AllowUnexported(item{})); diff != "" {
		t.Errorf("want items mismatch(-want +got):\n%s", diff)
	}

}
