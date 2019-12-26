package split

import "strings"

// split
// 将s按照sep分割，返回一个字符串的切片
// Split("i love you", "love") => ["i ", "you"]
func Split(s, sep string) (ret []string) {
	ret = make([]string, 0, strings.Count(s, sep)+1)
	idx := strings.Index(s, sep)
	for idx > -1 {
		ret = append(ret, s[:idx])
		s = s[idx+len(sep):]
		idx = strings.Index(s, sep)
	}
	ret = append(ret, s)
	return
}
