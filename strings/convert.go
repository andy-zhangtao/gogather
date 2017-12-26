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
