package util

import (
	"sort"
	"strings"
)

// buildsSort 对buildings进行排序
//Example:
// `{'123','456','234','567','345'}` 排序以后变为
//  {'123','234','345','456','567'}
func BuildsSort(s string) string {
	s = s[1 : len(s)-1]
	ss := strings.Split(s, `,`)
	sort.Strings(ss)
	s = strings.Join(ss, ",")
	s = `{` + s + `}`
	return s
}
