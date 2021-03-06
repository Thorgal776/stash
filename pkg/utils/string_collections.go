package utils

import "strconv"

// https://gobyexample.com/collection-functions

func StrIndex(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func StrInclude(vs []string, t string) bool {
	return StrIndex(vs, t) >= 0
}

func StrFilter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func StrMap(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

// StringSliceToIntSlice converts a slice of strings to a slice of ints. If any
// values cannot be parsed, then they are inserted into the returned slice as
// 0.
func StringSliceToIntSlice(ss []string) []int {
	ret := make([]int, len(ss))
	for i, v := range ss {
		ret[i], _ = strconv.Atoi(v)
	}

	return ret
}
