package strings

import (
	"strings"
	"testing"
)

func TestStdSymExstact(t *testing.T) {
	src := "{ABC}abc{DEF}jhi"
	sub, err := SymExstact(src, "{", "}")
	if err != nil {
		t.Error(err)
	}

	if len(sub) != 2 {
		t.Error("Exstact Error. Count Error")
	}

	if sub[0] != "ABC" || sub[1] != "DEF" {
		t.Errorf("Exstact Error. Value Error. 1:[%s]2:[%s]\n", sub[0], sub[1])
	}

}

func TestNestdSymExstact(t *testing.T) {
	src := "{{ABC}DEF}"
	sub, err := SymExstact(src, "{", "}")
	if err != nil {
		t.Error(err)
	}

	if len(sub) != 1 {
		t.Error("Exstact Error. Count Error")
	}

	if sub[0] != "{ABC}DEF" {
		t.Errorf("Exstact Error. Value Error. 1:[%s]\n", sub[0])
	}

	sub, err = SymExstact(sub[0], "{", "}")
	if err != nil {
		t.Error(err)
	}

	if len(sub) != 1 {
		t.Error("Exstact Error. Count Error")
	}

	if sub[0] != "ABC" {
		t.Errorf("Exstact Error. Value Error. 1:[%s]\n", sub[0])
	}
}

func TestExceptSymExstact(t *testing.T) {
	src := "{ABC"
	_, err := SymExstact(src, "{", "}")
	if err == nil {
		t.Error("{ABC Exstact Error")
	}

	if !strings.Contains(err.Error(), "does not contains") {
		t.Error("{ABC Exstact Error " + err.Error())
	}

	src = "{{ABC}"
	_, err = SymExstact(src, "{", "}")
	if err == nil {
		t.Error("{ABC Exstact Error")
	}

	if !strings.Contains(err.Error(), "The number of") {
		t.Error("{ABC Exstact Error " + err.Error())
	}

	src = "{{ABC}"
	_, err = SymExstact(src, "{", "{")
	if err == nil {
		t.Error("{ABC Exstact Error")
	}

	if !strings.Contains(err.Error(), "invoke DouExstact") {
		t.Error("{ABC Exstact Error " + err.Error())
	}
}
