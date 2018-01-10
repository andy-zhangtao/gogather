package strings

import (
	"errors"
	"strings"
	"regexp"
)

// SymExstact 对称截取. 从src中截取sym1与sym2之间的字符
// 例如从{ABC}abc{CHD}中按照'{','}'进行截取，最后会截取到
// ABC和CHD
func SymExstact(src string, sym1, sym2 string) ([]string, error) {
	if !strings.Contains(src, sym1) || !strings.Contains(src, sym2) {
		return nil, errors.New(src + " does not contains " + sym1 + " or " + sym2)
	}

	if strings.Count(src, sym1) != strings.Count(src, sym2) {
		return nil, errors.New("The number of " + sym1 + "," + sym2 + " is not equal")
	}

	if strings.Compare(sym1, sym2) == 0 {
		return nil, errors.New("Since " + sym1 + " is equal to " + sym2 + ", please invoke DouExstact")
	}

	var result []string
	isSym1 := true
	ids1 := strings.Index(src, sym1)
	sub := src[ids1+1:]

	for {
		ids2 := 0
		str := ""
		if !strings.Contains(sub, sym2) {
			break
		}

		if isSym1 {
			ids2 = strings.Index(sub, sym2)
			str = sub[:ids2]
			if strings.Contains(str, sym1) {
				count := strings.Count(str, sym1)
				ts := sub
				for index := 0; index < count; index++ {
					ids2 += strings.Index(ts, sym2)
					ts = ts[ids2:]
				}

				str = sub[:ids2]
			}

			sub = sub[ids2+1:]
			isSym1 = !isSym1

			result = append(result, str)
		} else {
			ids1 = strings.Index(sub, sym1)
			sub = sub[ids1+1:]

			isSym1 = !isSym1
		}

	}

	return result, nil
}

// DouExstact 标准截取. 从src中按照指定的sym进行截取sym之间的字符
// 例如从#ABC#abc#DEF#中截取出 ABC和DEF
func DouExstact(src string, sym string) ([]string, error) {
	count := strings.Count(src, sym)

	if count == 1 {
		return nil, errors.New("The src string must have more than one split char")
	}

	var result []string

	isSym := true
	ids1 := strings.Index(src, sym)
	sub := src[ids1+1:]

	for {
		if !strings.Contains(sub, sym) {
			break
		}

		ids2 := strings.Index(sub, sym)
		if isSym {
			str := sub[:ids2]
			sub = sub[ids2+1:]

			result = append(result, str)
			isSym = !isSym
		} else {
			sub = sub[ids2+1:]
			isSym = !isSym
		}

	}
	return result, nil

}


// RemoveMultipeSpace 去除字符串中多余的空格
func RemoveMultipeSpace(oldStr string) (newStr string) {
	re_leadclose_whtsp := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
	re_inside_whtsp := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	final := re_leadclose_whtsp.ReplaceAllString(oldStr, "")
	newStr = re_inside_whtsp.ReplaceAllString(final, " ")
	return
}
