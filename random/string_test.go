package random

import (
	"testing"
)

func TestGetString(t *testing.T) {
	randStr1 := GetString(6)
	if len(randStr1) != 6 {
		t.Error("Random String Generate Length Error!")
	}

	randStr2 := GetString(6)
	if len(randStr2) != 6 {
		t.Error("Random String Generate Length Error!")
	}

	if randStr1 == randStr2 {
		t.Error("String1 And String2 Equal")
	}
}
