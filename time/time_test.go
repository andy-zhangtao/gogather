package time

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

//Write by zhangtao<ztao8607@gmail.com> . In 2018/5/7.

func TestGetCurrentDayOfMonth(t *testing.T) {
	date, err := time.Parse("2006-01-02", "2016-01-30")
	assert.Empty(t, err, "Parse Time Error")
	day := GetCurrentDayOfMonth(date)
	assert.Equal(t, 30, day, "Get Day Of Month Error")
	date, err = time.Parse("2006-01-02", "2016-12-31")
	assert.Empty(t, err, "Parse Time Error")
	day = GetCurrentDayOfMonth(date)
	assert.Equal(t, 31, day, "Get Day Of Month Error")
}

func TestGetCurrentDayOfYear(t *testing.T) {
	date, err := time.Parse("2006-01-02", "2016-01-30")
	assert.Empty(t, err, "Get Day Of Year Error")
	day, err := GetCurrentDayOfYear(date)
	assert.Empty(t, err, "Get Day Of Year Error")
	assert.Equal(t, 30, day, "Get Day Of Year Error")
	date, err = time.Parse("2006-01-02", "2016-12-31")
	assert.Empty(t, err, "Get Day Of Year Error")
	day, err = GetCurrentDayOfYear(date)
	assert.Empty(t, err, "Get Day Of Year Error")
	assert.Equal(t, 366, day, "Get Day Of Year Error")
}
