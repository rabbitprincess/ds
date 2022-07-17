package binary_search

import (
	"sort"
)

func Sort_int(_n *[]int) {
	sort.Ints(*_n)
}

func Sort_float(_f8 *[]float64) {
	sort.Float64s(*_f8)
}

func Sort_string(_s *[]string) {
	sort.Strings(*_s)
}

func IsSorted_int(_n []int) (isSorted bool) {
	return sort.IntsAreSorted(_n)
}

func IsSorted_float(_f8 []float64) (isSorted bool) {
	return sort.Float64sAreSorted(_f8)
}

func IsSorted_string(_s []string) (isSorted bool) {
	return sort.StringsAreSorted(_s)
}

func GetPos_int(_n []int, _target int) (pos int) {
	pos = sort.SearchInts(_n, _target)
	if _n[pos] != _target {
		return -1
	}
	return pos
}

func GetPos_float(_f8 []float64, _target float64) (pos int) {
	pos = sort.SearchFloat64s(_f8, _target)
	if _f8[pos] != _target {
		return -1
	}
	return pos
}

func GetPos_string(_s []string, _target string) (pos int) {
	pos = sort.SearchStrings(_s, _target)
	if _s[pos] != _target {
		return -1
	}
	return pos
}
