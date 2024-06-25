package makesquare

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestMakeSquare_NoSolutions(t *testing.T) {
	tt := [][]int{
		{3, 3, 3, 4},
		{15, 5, 10, 10},
		{},
		{1},
		{1, 2},
		{1, 2, 3},
	}

	for i := range tt {
		if makesquare(tt[i]) {
			t.Errorf("некорректный результат для %v", tt[i])
		}
	}
}

func TestGetPositions(t *testing.T) {
	tt := map[string]struct {
		m    int
		side int
		want [][]int
	}{
		"1": {
			m:    1,
			side: 3,
			want: [][]int{
				{0},
				{1},
				{2},
			},
		},
		"2": {
			m:    2,
			side: 5,
			want: [][]int{
				{0, 1},
				{1, 2},
				{2, 3},
				{3, 4},
			},
		},
		"3": {
			m:    3,
			side: 3,
			want: [][]int{
				{0, 1, 2},
			},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := getPositions(tc.m, tc.side)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("некорректный результат (-want +got):\n%s", diff)
			}
		})
	}
}

func TestMakeSquare(t *testing.T) {
	tt := map[string]struct {
		in   []int
		want bool
	}{
		"1": {
			in:   []int{1, 1, 2, 2, 2},
			want: true,
		},
		"2": {
			in:   []int{3, 3, 3, 3, 4},
			want: false,
		},
		"3": {
			in:   []int{2, 2, 2, 2, 2, 8, 8, 4, 10},
			want: true,
		},
		"4": {
			in:   []int{3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 6},
			want: false,
		},
		"5": {
			in:   []int{20, 13, 19, 19, 4, 15, 10, 5, 5, 15, 14, 11, 3, 20, 11},
			want: true,
		},
		"6": {
			in:   []int{3, 1, 3, 3, 10, 7, 10, 3, 6, 9, 10, 3, 7, 6, 7},
			want: true,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			if got := makesquare(tc.in); got != tc.want {
				t.Errorf("failed on %v; expected %t", tc.in, tc.want)
			}
		})
	}
}
