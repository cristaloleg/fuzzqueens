package fuzzqueens

import (
	"testing"
)

type pos struct {
	// NOTE: no X because X is an index [0,7]
	y, sum, diff int
}

func FuzzQueens(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		if len(data) != 8 {
			t.Skip()
			return
		}

		for _, d := range data {
			if d < 0 || d > 7 {
				t.Skip()
				return
			}
		}

		var queens [8]pos
		for x, d := range data {
			y := int(d)
			queens[x] = pos{
				y:    y,
				sum:  x + y,
				diff: x - y,
			}
		}

		if areConnected(queens) {
			return
		}

		t.Errorf("solution: (%d,%d), (%d, %d), (%d, %d), (%d, %d), (%d, %d), (%d, %d), (%d, %d), (%d, %d)",
			0, queens[0].y,
			1, queens[1].y,
			2, queens[2].y,
			3, queens[3].y,
			4, queens[4].y,
			5, queens[5].y,
			6, queens[6].y,
			7, queens[7].y,
		)
	})
}

// no loops, just for fun
func areConnected(q [8]pos) bool {
	switch {
	// 1
	case q[1].y == q[0].y,
		q[1].sum == q[0].sum,
		q[1].diff == q[0].diff:
		return true

		// 2
	case q[2].y == q[0].y,
		q[2].sum == q[0].sum,
		q[2].diff == q[0].diff:
		return true
	case q[2].y == q[1].y,
		q[2].sum == q[1].sum,
		q[2].diff == q[1].diff:
		return true

		// 3
	case q[3].y == q[0].y,
		q[3].sum == q[0].sum,
		q[3].diff == q[0].diff:
		return true
	case q[3].y == q[1].y,
		q[3].sum == q[1].sum,
		q[3].diff == q[1].diff:
		return true
	case q[3].y == q[2].y,
		q[3].sum == q[2].sum,
		q[3].diff == q[2].diff:
		return true

		// 4
	case q[4].y == q[0].y,
		q[4].sum == q[0].sum,
		q[4].diff == q[0].diff:
		return true
	case q[4].y == q[1].y,
		q[4].sum == q[1].sum,
		q[4].diff == q[1].diff:
		return true
	case q[4].y == q[2].y,
		q[4].sum == q[2].sum,
		q[4].diff == q[2].diff:
		return true
	case q[4].y == q[3].y,
		q[4].sum == q[3].sum,
		q[4].diff == q[3].diff:
		return true
		// 5
	case q[5].y == q[0].y,
		q[5].sum == q[0].sum,
		q[5].diff == q[0].diff:
		return true
	case q[5].y == q[1].y,
		q[5].sum == q[1].sum,
		q[5].diff == q[1].diff:
		return true
	case q[5].y == q[2].y,
		q[5].sum == q[2].sum,
		q[5].diff == q[2].diff:
		return true
	case q[5].y == q[3].y,
		q[5].sum == q[3].sum,
		q[5].diff == q[3].diff:
		return true
	case q[5].y == q[4].y,
		q[5].sum == q[4].sum,
		q[5].diff == q[4].diff:
		return true
		// 6
	case q[6].y == q[0].y,
		q[6].sum == q[0].sum,
		q[6].diff == q[0].diff:
		return true
	case q[6].y == q[1].y,
		q[6].sum == q[1].sum,
		q[6].diff == q[1].diff:
		return true
	case q[6].y == q[2].y,
		q[6].sum == q[2].sum,
		q[6].diff == q[2].diff:
		return true
	case q[6].y == q[3].y,
		q[6].sum == q[3].sum,
		q[6].diff == q[3].diff:
		return true
	case q[6].y == q[4].y,
		q[6].sum == q[4].sum,
		q[6].diff == q[4].diff:
		return true
	case q[6].y == q[5].y,
		q[6].sum == q[5].sum,
		q[6].diff == q[5].diff:
		return true

		// 7
	case q[7].y == q[0].y,
		q[7].sum == q[0].sum,
		q[7].diff == q[0].diff:
		return true
	case q[7].y == q[1].y,
		q[7].sum == q[1].sum,
		q[7].diff == q[1].diff:
		return true
	case q[7].y == q[2].y,
		q[7].sum == q[2].sum,
		q[7].diff == q[2].diff:
		return true
	case q[7].y == q[3].y,
		q[7].sum == q[3].sum,
		q[7].diff == q[3].diff:
		return true
	case q[7].y == q[4].y,
		q[7].sum == q[4].sum,
		q[7].diff == q[4].diff:
		return true
	case q[7].y == q[5].y,
		q[7].sum == q[5].sum,
		q[7].diff == q[5].diff:
		return true
	case q[7].y == q[6].y,
		q[7].sum == q[6].sum,
		q[7].diff == q[6].diff:
		return true

	default:
		return false
	}
}
