package nqueens

import (
	"dlx"
	"testing"
)

func TestNQueens(t *testing.T) {
	tt := map[string]struct {
		n    int
		want int
	}{
		"2": {
			n:    2,
			want: 0,
		},
		"3": {
			n:    3,
			want: 0,
		},
		"4": {
			n:    4,
			want: 2,
		},
		"5": {
			n:    5,
			want: 10,
		},
		"6": {
			n:    6,
			want: 4,
		},
		"7": {
			n:    7,
			want: 40,
		},
		"8": {
			n:    8,
			want: 92,
		},
		"9": {
			n:    9,
			want: 352,
		},
		"10": {
			n:    10,
			want: 724,
		},
		"11": {
			n:    11,
			want: 2_680,
		},
		"12": {
			n:    12,
			want: 14_200,
		},
		"13": {
			n:    13,
			want: 73_712,
		},
		"14": {
			n:    14,
			want: 365_596,
		},
	}

	for name, tc := range tt {
		// https://go.dev/blog/subtests
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			m := prepareMatrix(tc.n)

			d, err := dlx.New(m)
			if err != nil {
				t.Fatal(err)
			}

			solutions, err := d.SolveString()
			if err != nil {
				t.Fatal(err)
			}

			if len(solutions) != tc.want {
				t.Fatalf("got %d solutions instead of %d", len(solutions), tc.want)
			}

			for i := 0; i < min(2, len(solutions)); i++ {
				printBoard(solutions[i], tc.n)
			}
		})
	}

}
