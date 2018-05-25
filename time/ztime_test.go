package time

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestZtime(t *testing.T) {

	zt := new(Ztime)

	tt, err := zt.Now().SetLocation("Asia/Shanghai").Format("2006-01-02T15:04")
	assert.Nil(t, err)

	l, _ := time.LoadLocation("Asia/Shanghai")

	assert.Equal(t, time.Now().In(l).Format("2006-01-02T15:04"), tt)

	tt, err = zt.Now().SetLocation("Asia/Shanghai/1").Format(time.RFC3339)
	assert.Error(t, err)

	tt, err = new(Ztime).Now().UTC().AddHour(7).Format("2006-01-02T15:04")

	assert.Equal(t, time.Now().In(l).Add(-1 * time.Hour).Format("2006-01-02T15:04"), tt)

	tt, err = new(Ztime).Now().UTC().AddHour(7).Format("YYYY-MM-DDThh:mm")

	assert.Equal(t, time.Now().In(l).Add(-1 * time.Hour).Format("2006-01-02T15:04"), tt)

	tt, err = new(Ztime).Now().UTC().AddHour(7).Format("YYYY-MM-DD hh:mm")

	assert.Equal(t, time.Now().In(l).Add(-1 * time.Hour).Format("2006-01-02 15:04"), tt)
}
