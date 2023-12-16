package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	s := "aab"

	r := arrangements(s)
	fmt.Println(r)
}

func arrangements(str string) string {
	tele := make([]int, 26)

	for _, char := range str {
		if char >= 'A' && char <= 'Z' {
			return ""
		}
		// [0] => a , [1] => b ...
		diff := char - 'a'
		tele[diff]++
	}

	// calculate maximum appearing char
	mx := slices.Max(tele)
	n := len(str)

	/// if max is greater than the half of the given string then its logicaly impossible to rearrange the string
	/// with the given conditions
	if mx > (n+1)/2 {
		return ""
	}

	/// generate pair of [char , appearing]
	pair := [][]int{}
	for i, v := range tele {
		if v > 0 {
			pair = append(pair, []int{v, i})
		}
	}

	// sort the pairs according to their appearings
	sort.Slice(pair, func(i, j int) bool {
		return pair[i][0] > pair[j][0]
	})

	ans := make([]byte, n)
	k := 0
	for _, e := range pair {
		v, i := e[0], e[1]
		for v > 0 {
			ans[k] = byte('a' + i)
			k += 2
			if k >= n {
				k = 1
			}
			v--
		}
	}
	return string(ans)
}
