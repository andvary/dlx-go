package mydlx

import (
	"errors"
	"github.com/google/go-cmp/cmp"
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
	d := &DLX{}

	if err := d.readInput("./testdata/input_good_3.txt"); err != nil {
		t.Fatal(err)
	}

	want := &DLX{
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
				items: map[int]struct{}{3: {}, 5: {}, 6: {}},
			},
			{
				prev:  0,
				next:  3,
				items: map[int]struct{}{1: {}, 4: {}, 7: {}},
			},
			{
				prev:  2,
				next:  4,
				items: map[int]struct{}{2: {}, 3: {}, 6: {}},
			},
			{
				prev:  3,
				next:  5,
				items: map[int]struct{}{1: {}, 4: {}},
			},
			{
				prev:  4,
				next:  6,
				items: map[int]struct{}{2: {}, 7: {}},
			},
			{
				prev:  5,
				next:  0,
				items: map[int]struct{}{4: {}, 5: {}, 7: {}},
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
	d := &DLX{}

	if err := d.readInput("./testdata/input_good_3.txt"); err != nil {
		t.Fatal(err)
	}

	want := &DLX{
		items: []*item{
			{name: "", prev: 6, next: 2},          //0
			{name: "a", prev: 0, next: 2, cnt: 0}, //1 x
			{name: "b", prev: 0, next: 3, cnt: 1}, //2
			{name: "c", prev: 2, next: 5, cnt: 2}, //3
			{name: "d", prev: 3, next: 5, cnt: 0}, //4 x
			{name: "e", prev: 3, next: 6, cnt: 1}, //5
			{name: "f", prev: 5, next: 0, cnt: 2}, //6
			{name: "g", prev: 6, next: 0, cnt: 0}, //7 x
		},
		opts: []*opt{
			{ //0
				prev:  3,
				next:  1,
				items: nil,
			},
			{ //1
				prev:  0,
				next:  3,
				items: map[int]struct{}{3: {}, 5: {}, 6: {}},
			},
			{ //2
				prev:  1,
				next:  3,
				items: map[int]struct{}{1: {}, 4: {}, 7: {}}, //x
			},
			{ //3
				prev:  1,
				next:  0,
				items: map[int]struct{}{2: {}, 3: {}, 6: {}},
			},
			{ //4
				prev:  3,
				next:  5,
				items: map[int]struct{}{1: {}, 4: {}}, //x
			},
			{ //5
				prev:  4,
				next:  6,
				items: map[int]struct{}{2: {}, 7: {}}, //x
			},
			{ //6
				prev:  5,
				next:  0,
				items: map[int]struct{}{4: {}, 5: {}, 7: {}}, //x
			},
		},
	}

	if err := d.cover(1); err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(want.items, d.items, cmp.AllowUnexported(item{})); diff != "" {
		t.Errorf("want items mismatch(-want +got):\n%s", diff)
	}

	// Сравниваем только состояние неудалённых опций, т.к. удаляются они не всегда в одном и том же порядке,
	// соответственно, указатели в них могут отличаться от теста к тесту.
	for i := d.opts[0].next; i != 0; i = d.opts[i].next {
		if diff := cmp.Diff(want.opts[i], d.opts[i], cmp.AllowUnexported(opt{})); diff != "" {
			t.Errorf("want options mismatch(-want +got):\n%s", diff)
		}
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
			{name: "", prev: 5, next: 2},          //0
			{name: "a", prev: 0, next: 2, cnt: 2}, //1	x
			{name: "b", prev: 0, next: 3, cnt: 2}, //2
			{name: "c", prev: 2, next: 4, cnt: 2}, //3
			{name: "d", prev: 3, next: 5, cnt: 3}, //4
			{name: "e", prev: 4, next: 0, cnt: 2}, //5
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

	removedItems := []int{1, 7, 6}
	d.restoreItems(removedItems)

	if diff := cmp.Diff(want, d.items, cmp.AllowUnexported(item{})); diff != "" {
		t.Errorf("want items mismatch(-want +got):\n%s", diff)
	}

}
