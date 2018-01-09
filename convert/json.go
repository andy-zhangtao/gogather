package convert

import (
	"errors"
)

func FindJsonFromStr(oldStr string) ([]string, error) {
	var newStr []string

	str := oldStr
	for {

		tmpStr, idx, err := splitJsonBySep(str, str[0:1])
		if err != nil {
			return newStr, err
		}

		newStr = append(newStr, tmpStr)
		str = str[idx+1:]
		if len(str) == 0 {
			break
		}
	}

	return newStr, nil
}

// splitJsonBySep 按照指定的起始符来查找合法的Json字符串
// sep是字符串起始字符，目前必须为'{'或者'['
// 返回的是查找到的json字符，同时返回当前字符位置
func splitJsonBySep(oldStr, sep string) (newStr string, idx int, err error) {
	endSep := ""
	switch sep {
	case "{":
		endSep = "}"
	case "[":
		endSep = "]"
	}

	if endSep == "" {
		err = errors.New("Begin Sep Invaild. Must Be '{' or ']'")
		return
	}

	// ptr用作层级标示 当ptr再次为0的时候，就是找到完整json的时候
	ptr := 0
	for i, s := range oldStr {
		newStr += string(s)
		if string(s) == sep {
			ptr++
		}
		if string(s) == endSep {
			ptr--
		}

		if ptr == 0 {
			idx = i
			break
		}
	}

	return
}
