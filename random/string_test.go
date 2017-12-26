package random

import (
	"testing"
	"strconv"
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

func TestGetRandom(t *testing.T) {
	randNum1 := GetRandom(10)
	if len(randNum1) != 10{
		t.Error("Random String Generate Length Error!")
	}

	if _, err := strconv.Atoi(randNum1); err != nil{
		t.Error("Random String Generate Number Error!")
	}

	randNum2 := GetRandom(6)
	if len(randNum2) != 6{
		t.Error("Random String Generate Length Error!")
	}

	if _, err := strconv.Atoi(randNum2); err != nil{
		t.Error("Random String Generate Number Error!")
	}
}