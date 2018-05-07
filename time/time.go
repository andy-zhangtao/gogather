package time

import (
	"time"
	"strconv"
	"strings"
	"errors"
	"fmt"
)

// GetTimeStamp 获取当前时间戳
// length 10:获取秒 13:获取毫秒
func GetTimeStamp(length int) string {
	return strconv.FormatInt(time.Now().UTC().UnixNano(), 10)[:length]
}

// GetCurrentDayOfMonth 获取给定的时间是当月的第几天
func GetCurrentDayOfMonth(t time.Time) (day int) {
	_, _, day = t.Date()
	return
}

func GetCurrentDayOfYear(t time.Time) (day int, err error) {
	current := t.Format("2006-01-02")
	cs := strings.Split(current, "-")

	if len(cs) != 3 {
		err = errors.New(fmt.Sprintf("Wrong Time [%v]", t))
		return
	}

	year, err := strconv.Atoi(cs[0])
	if err != nil {
		return
	}
	month, err := strconv.Atoi(cs[1])
	if err != nil {
		return
	}

	d, err := strconv.Atoi(cs[2])
	if err != nil {
		return
	}

	var num int
	switch month {
	case 1:
		num = 0
	case 2:
		num = 31
	case 3:
		if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
			num = 31 + 29
		} else {
			num = 31 + 28
		}
	case 4:
		if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
			num = 31 + 29 + 31
		} else {
			num = 31 + 28 + 31
		}
	case 5:
		if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
			num = 31 + 29 + 31 + 30
		} else {
			num = 31 + 28 + 31 + 30
		}
	case 6:
		if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
			num = 31 + 29 + 31 + 30 + 31
		} else {
			num = 31 + 28 + 31 + 30 + 31
		}
	case 7:
		if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
			num = 31 + 29 + 31 + 30 + 31 + 30
		} else {
			num = 31 + 28 + 31 + 30 + 31 + 30
		}
	case 8:
		if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
			num = 31 + 29 + 31 + 30 + 31 + 30 + 31
		} else {
			num = 31 + 28 + 31 + 30 + 31 + 30 + 31
		}
	case 9:
		if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
			num = 31 + 29 + 31 + 30 + 31 + 30 + 31 + 31
		} else {
			num = 31 + 28 + 31 + 30 + 31 + 30 + 31 + 31
		}
	case 10:
		if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
			num = 31 + 29 + 31 + 30 + 31 + 30 + 31 + 31 + 30
		} else {
			num = 31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30
		}
	case 11:
		if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
			num = 31 + 29 + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31
		} else {
			num = 31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31
		}
	case 12:
		if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
			num = 31 + 29 + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31 + 30
		} else {
			num = 31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31 + 30
		}
	}

	return d + num, nil
}
