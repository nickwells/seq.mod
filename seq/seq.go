package seq

import "math"

// absInt64 returns the absolute value of the argument
func absInt64(i int64) int64 {
	if i < 0 {
		i *= -1
	}
	return i
}

// absInt returns the absolute value of the argument
func absInt(i int) int {
	if i < 0 {
		i *= -1
	}
	return i
}

// NormaliseIncrInt64 ensures that the sign of incr is such that repeatedly
// adding it to first will take you ever closer to last.
func NormaliseIncrInt64(first, last, incr int64) int64 {
	if (last < first && incr > 0) ||
		(last > first && incr < 0) {
		incr *= -1
	}
	return incr
}

// NormaliseIncrFloat64 ensures that the sign of incr is such that repeatedly
// adding it to first will take you ever closer to last.
func NormaliseIncrFloat64(first, last, incr float64) float64 {
	if (last < first && incr > 0) ||
		(last > first && incr < 0) {
		incr *= -1
	}
	return incr
}

// Int64 returns a slice of int64 values starting at first and going up by
// incr until the final entry is greater than or equal to last. If last is
// less than first then subsequent entries will go down by incr.
//
// If incr is zero then an empty slice is returned. Otherwise the sign of
// incr is ignored and will be automatically adjusted such that repeatedly
// adding incr to first will get you ever closer to last.
func Int64(first, last, incr int64) []int64 {
	if incr == 0 {
		return []int64{}
	}
	if first == last {
		return []int64{first}
	}

	incr = NormaliseIncrInt64(first, last, incr)

	// by pre-calculating the slice size we
	// (a) avoid any problems of arithmetic overflow with large increments
	//     towards a limit that is close to the maximum value of an int64
	// we also
	// (b) reduce the number of memory allocations (a beneficial side effect)
	count := int(1 + (absInt64(last-first) / absInt64(incr)))
	r := make([]int64, 0, count)

	val := first
	for i := 0; i < count; i++ {
		r = append(r, val)
		val += incr
	}
	return r
}

// Float64Dup returns a slice of count copies of val
func Float64Dup(val float64, count int) []float64 {
	count = absInt(count)
	if count == 0 {
		return []float64{}
	}
	if val == 0.0 {
		return make([]float64, count)
	}

	r := make([]float64, 0, count)
	for i := 0; i < count; i++ {
		r = append(r, val)
	}
	return r
}

// Int64Dup returns a slice of count copies of val
func Int64Dup(val int64, count int) []int64 {
	count = absInt(count)
	if count == 0 {
		return []int64{}
	}

	r := make([]int64, 0, count)

	for i := 0; i < count; i++ {
		r = append(r, val)
	}
	return r
}

// Int64ByLen returns a slice of int64 values starting at first and finishing
// at last having count entries. If count equals zero an empty slice is
// returned, if it's one then a slice containing just the first entry is
// returned.
func Int64ByLen(first, last int64, count int) []int64 {
	count = absInt(count)
	switch count {
	case 0:
		return []int64{}
	case 1:
		return []int64{first}
	case 2:
		return []int64{first, last}
	}

	if first == last {
		return Int64Dup(first, count)
	}

	r := make([]int64, 0, count)
	interval := absInt64(last - first)
	var incr float64 = NormaliseIncrFloat64(float64(first), float64(last),
		float64(interval)/float64(count-1))

	val := float64(first)
	for i := 0; i < count; i++ {
		r = append(r, int64(math.Round(val)))
		val += incr
	}
	return r
}

// Float64 returns a slice of float64 values starting at first and going up
// by incr until the final entry is greater than or equal to last. If last is
// less than first then subsequent entries will go down by incr.
//
// If incr is zero then an empty slice is returned. Otherwise the sign of
// incr is ignored and will be automatically adjusted such that repeatedly
// adding incr to first will get you ever closer to last.
func Float64(first, last, incr float64) []float64 {
	if incr == 0 {
		return []float64{}
	}
	if first == last {
		return []float64{first}
	}

	incr = NormaliseIncrFloat64(first, last, incr)

	count := int(1 + absInt64(int64((last-first)/incr)))
	r := make([]float64, count)

	val := first
	for i := 0; i < count/2; i++ {
		r[i] = val
		val += incr
	}
	val = last
	for i := count - 1; i >= count/2; i-- {
		r[i] = val
		val -= incr
	}
	return r
}
