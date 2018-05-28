package zlog

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestZlog_Error(t *testing.T) {
	z := GetZlog()

	z.AddID("a_custom_track_id")

	err := z.Error("sample error")

	assert.EqualError(t, err, "_track [a_custom_track_id] sample error")
}

func TestZlog_MyTrack(t *testing.T) {
	z := GetZlog()

	id := z.MyTrack()

	assert.Equal(t, 12, len(id))
}
