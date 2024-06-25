package dlx

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestNewBitvector(t *testing.T) {
	tt := map[string]struct {
		n    int
		want int
	}{
		"1": {
			n:    0,
			want: 1,
		},
		"2": {
			n:    64,
			want: 2,
		},
		"3": {
			n:    127,
			want: 2,
		},
		"4": {
			n:    5999,
			want: 94,
		},
		"5": {
			n:    128,
			want: 3,
		},
		"6": {
			n:    63,
			want: 1,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			b := newBitvector(tc.n)
			if len(b.v) != tc.want {
				t.Errorf("для %d элементов ожидали длину %d, получили %d", tc.n, tc.want, len(b.v))
			}
		})
	}
}

func TestBitvectorAdd(t *testing.T) {
	in := []int{0, 1, 6, 63, 64, 65, 127, 128}
	want := make([]uint64, 3)
	want[0] = want[0] | 1 | 1<<1 | 1<<6 | 1<<63
	want[1] = want[1] | 1 | 1<<1 | 1<<63
	want[2] = want[2] | 1

	b := newBitvector(128)
	for i := range in {
		b.add(in[i])

	}

	if diff := cmp.Diff(want, b.v); diff != "" {
		t.Errorf("неожиданный результат(-want +got):\n%s", diff)
		t.Logf("%b %b %b\n", b.v[0], b.v[1], b.v[2])
	}

}

func TestBitvectorIsPresent(t *testing.T) {
	tt := map[string]struct {
		in   int
		want bool
	}{
		"1": {
			in:   0,
			want: true,
		},
		"2": {
			in:   1,
			want: true,
		},
		"3": {
			in:   6,
			want: true,
		},
		"4": {
			in:   63,
			want: true,
		},
		"5": {
			in:   64,
			want: true,
		},
		"6": {
			in:   65,
			want: true,
		},
		"7": {
			in:   127,
			want: true,
		},
		"8": {
			in:   128,
			want: true,
		},
		"9": {
			in:   2,
			want: false,
		},
		"10": {
			in:   129,
			want: false,
		},
		"11": {
			in:   66,
			want: false,
		},
	}

	in := []int{0, 1, 6, 63, 64, 65, 127, 128}
	b := newBitvector(128)
	for i := range in {
		b.add(in[i])

	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			if got := b.isPresent(tc.in); got != tc.want {
				t.Errorf("неожиданный результат для %d; ожидали %t", tc.in, tc.want)
			}
		})
	}

}
