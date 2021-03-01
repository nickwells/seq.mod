package seq_test

import (
	"fmt"
	"testing"

	"github.com/nickwells/seq.mod/seq"
	"github.com/nickwells/testhelper.mod/testhelper"
)

func TestInt64(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		first, last, incr int64
		expSeq            []int64
	}{
		{
			ID:     testhelper.MkID("crossing zero"),
			first:  -1,
			last:   1,
			incr:   1,
			expSeq: []int64{-1, 0, 1},
		},
		{
			ID:     testhelper.MkID("crossing zero, backwards"),
			first:  1,
			last:   -1,
			incr:   1,
			expSeq: []int64{1, 0, -1},
		},
		{
			ID:     testhelper.MkID("crossing zero, non-unit incr"),
			first:  -1,
			last:   1,
			incr:   2,
			expSeq: []int64{-1, 1},
		},
		{
			ID:     testhelper.MkID("first==last"),
			first:  1,
			last:   1,
			incr:   1,
			expSeq: []int64{1},
		},
		{
			ID:     testhelper.MkID("first<last"),
			first:  1,
			last:   3,
			incr:   1,
			expSeq: []int64{1, 2, 3},
		},
		{
			ID:     testhelper.MkID("first>last"),
			first:  5,
			last:   3,
			incr:   1,
			expSeq: []int64{5, 4, 3},
		},
		{
			ID:     testhelper.MkID("first<last, big incr"),
			first:  1,
			last:   3,
			incr:   10,
			expSeq: []int64{1},
		},
		{
			ID:     testhelper.MkID("first>last, big incr"),
			first:  5,
			last:   3,
			incr:   10,
			expSeq: []int64{5},
		},
		{
			ID:     testhelper.MkID("first>last, negative"),
			first:  -1,
			last:   -3,
			incr:   1,
			expSeq: []int64{-1, -2, -3},
		},
		{
			ID:     testhelper.MkID("first<last, negative"),
			first:  -5,
			last:   -3,
			incr:   1,
			expSeq: []int64{-5, -4, -3},
		},
		{
			ID:     testhelper.MkID("first>last, negative, big incr"),
			first:  -1,
			last:   -3,
			incr:   10,
			expSeq: []int64{-1},
		},
		{
			ID:     testhelper.MkID("first<last, negative, big incr"),
			first:  -5,
			last:   -3,
			incr:   10,
			expSeq: []int64{-5},
		},
		{
			ID:     testhelper.MkID("incr==0"),
			first:  5,
			last:   3,
			incr:   0,
			expSeq: []int64{},
		},
		{
			ID:     testhelper.MkID("very large values"),
			first:  0,
			last:   2000000000,
			incr:   1999999999,
			expSeq: []int64{0, 1999999999},
		},
		{
			ID:     testhelper.MkID("very large values (backwards)"),
			first:  0,
			last:   -2000000000,
			incr:   1999999999,
			expSeq: []int64{0, -1999999999},
		},
	}

	for _, tc := range testCases {
		r := seq.Int64(tc.first, tc.last, tc.incr)
		testhelper.DiffInt64Slice(t, tc.IDStr(), "slice", r, tc.expSeq)
	}
}

func TestInt64ByLen(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		first, last int64
		count       int
		expSeq      []int64
	}{
		{
			ID:     testhelper.MkID("count: 0"),
			first:  1,
			last:   3,
			count:  0,
			expSeq: []int64{},
		},
		{
			ID:     testhelper.MkID("count: 1"),
			first:  1,
			last:   3,
			count:  1,
			expSeq: []int64{1},
		},
		{
			ID:     testhelper.MkID("count: 2"),
			first:  1,
			last:   3,
			count:  2,
			expSeq: []int64{1, 3},
		},
		{
			ID:     testhelper.MkID("count: 3"),
			first:  1,
			last:   3,
			count:  3,
			expSeq: []int64{1, 2, 3},
		},
		{
			ID:     testhelper.MkID("count: 7"),
			first:  1,
			last:   3,
			count:  7,
			expSeq: []int64{1, 1, 2, 2, 2, 3, 3},
		},
	}

	for _, tc := range testCases {
		r := seq.Int64ByLen(tc.first, tc.last, tc.count)
		testhelper.DiffInt64Slice(t, tc.IDStr(), "slice", r, tc.expSeq)
	}
}

func TestFloat64(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		first, last, incr float64
		epsilon           float64
		hasExpSeq         bool
		expSeq            []float64
		expLen            int
	}{
		{
			ID:        testhelper.MkID("zero incr"),
			first:     1.5,
			last:      2.5,
			incr:      0,
			epsilon:   0,
			hasExpSeq: true,
			expSeq:    []float64{},
		},
		{
			ID:        testhelper.MkID("non-zero incr, first==last"),
			first:     1.5,
			last:      1.5,
			incr:      0.1,
			epsilon:   0,
			hasExpSeq: true,
			expSeq:    []float64{1.5},
		},
		{
			ID:        testhelper.MkID("non-zero incr, first!=last"),
			first:     1.5,
			last:      2.0,
			incr:      0.1,
			epsilon:   0.0000001,
			hasExpSeq: true,
			expSeq:    []float64{1.5, 1.6, 1.7, 1.8, 1.9, 2.0},
		},
		{
			ID:      testhelper.MkID("small incr, first!=last"),
			first:   1.5,
			last:    2.5,
			incr:    0.000001,
			epsilon: 0.0000001,
			expLen:  1000001,
		},
		{
			ID:      testhelper.MkID("small incr, first!=last, both -ve"),
			first:   -1.5,
			last:    -2.5,
			incr:    0.000001,
			epsilon: 0.0000001,
			expLen:  1000001,
		},
	}

	for _, tc := range testCases {
		id := tc.IDStr()
		r := seq.Float64(tc.first, tc.last, tc.incr)
		if tc.hasExpSeq {
			if testhelper.DiffFloat64Slice(t, id, "slice",
				r, tc.expSeq, tc.epsilon) {
				continue
			}
			tc.expLen = len(tc.expSeq)
		}
		if testhelper.DiffInt(t, id, "slice len", len(r), tc.expLen) {
			continue
		}
		if len(r) > 0 {
			if testhelper.DiffFloat64(t, id, "1st", r[0], tc.first, 0) ||
				testhelper.DiffFloat64(t, id, "last", r[len(r)-1], tc.last, 0) {
				continue
			}

			incr := seq.NormaliseIncrFloat64(tc.first, tc.last, tc.incr)
			for i := 1; i < len(r); i++ {
				diff := r[i] - r[i-1]
				name := fmt.Sprintf("increment (%d-%d)", i-1, i)
				if testhelper.DiffFloat64(t, id, name, diff, incr, tc.epsilon) {
					continue
				}
			}
		}
	}
}

func TestDupInt64(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		val    int64
		count  int
		expSeq []int64
	}{
		{
			ID:     testhelper.MkID("count == 0"),
			val:    1,
			count:  0,
			expSeq: []int64{},
		},
		{
			ID:     testhelper.MkID("count == 1, val == 0"),
			val:    0,
			count:  1,
			expSeq: []int64{0},
		},
		{
			ID:     testhelper.MkID("count == 2, val == 0"),
			val:    0,
			count:  2,
			expSeq: []int64{0, 0},
		},
		{
			ID:     testhelper.MkID("count == -2, val == 0"),
			val:    0,
			count:  -2,
			expSeq: []int64{0, 0},
		},
		{
			ID:     testhelper.MkID("count == 1, val == 42"),
			val:    42,
			count:  1,
			expSeq: []int64{42},
		},
		{
			ID:     testhelper.MkID("count == 2, val == 42"),
			val:    42,
			count:  2,
			expSeq: []int64{42, 42},
		},
		{
			ID:     testhelper.MkID("count == -2, val == 42"),
			val:    42,
			count:  -2,
			expSeq: []int64{42, 42},
		},
	}

	for _, tc := range testCases {
		r := seq.Int64Dup(tc.val, tc.count)
		testhelper.DiffInt64Slice(t, tc.IDStr(), "slice", r, tc.expSeq)
	}
}

func TestDupFloat64(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		val    float64
		count  int
		expSeq []float64
	}{
		{
			ID:     testhelper.MkID("count == 0"),
			val:    1,
			count:  0,
			expSeq: []float64{},
		},
		{
			ID:     testhelper.MkID("count == 1, val == 0"),
			val:    0,
			count:  1,
			expSeq: []float64{0},
		},
		{
			ID:     testhelper.MkID("count == 2, val == 0"),
			val:    0,
			count:  2,
			expSeq: []float64{0, 0},
		},
		{
			ID:     testhelper.MkID("count == -2, val == 0"),
			val:    0,
			count:  -2,
			expSeq: []float64{0, 0},
		},
		{
			ID:     testhelper.MkID("count == 1, val == 42"),
			val:    42,
			count:  1,
			expSeq: []float64{42},
		},
		{
			ID:     testhelper.MkID("count == 2, val == 42"),
			val:    42,
			count:  2,
			expSeq: []float64{42, 42},
		},
		{
			ID:     testhelper.MkID("count == -2, val == 42"),
			val:    42,
			count:  -2,
			expSeq: []float64{42, 42},
		},
	}

	for _, tc := range testCases {
		r := seq.Float64Dup(tc.val, tc.count)
		testhelper.DiffFloat64Slice(t, tc.IDStr(), "slice", r, tc.expSeq, 0)
	}
}
