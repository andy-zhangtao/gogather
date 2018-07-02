package strings

import "strings"

// Reverse 将源字符串按照指定的分隔符进行反转
// 例如:
// src = www.mydomain.com.cn
// dist = cn.com.mydomian.www
func Reverse(src, seg string) string {
	s := strings.Split(src, seg)
	i := len(s)
	dist := make([]string, i)
	for _, str := range s {
		i--
		dist[i] = str
	}

	return strings.Join(dist, seg)
}

// ReverseWithSeg 将源字符串按照指定的分隔符进行反转,同时使用新分隔符替代旧分隔符
// 例如:
// src = www.mydomain.com.cn, newSeg="-"
// dist = cn-com-mydomian-www
func ReverseWithSeg(src, oldSeg, newSeg string) string {
	return strings.Replace(Reverse(src, oldSeg), oldSeg, newSeg, -1)
}

// ReplaceAscii 替换src当中指定的字符串类型的控制字符
// 例如将 字符串类型的"\n"替换成 Ascii的 10(LF)
func ReplaceAscii(src string, seg []string) string {

	for _, s := range seg {
		switch s {
		case "\\n":
			src = strings.Replace(src, "\\n", "\n", -1)
		case "\\t":
			src = strings.Replace(src, "\\t", "\t", -1)
		}
	}

	return src
}
