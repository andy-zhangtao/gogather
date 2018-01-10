package strings

import (
	"log"
	"strings"
	"testing"
	"github.com/stretchr/testify/assert"
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

func TestDouExstact(t *testing.T) {
	src := "#ABC#DEF#HHH#eee"

	sub, err := DouExstact(src, "#")
	if err != nil {
		t.Error(err)
	}

	if len(sub) != 2 {
		t.Error("Exstact Error. Count Error")
	}

	if sub[0] != "ABC" || sub[1] != "HHH" {
		t.Errorf("Exstact Error. Value Error. 1:[%s]2:[%s]\n", sub[0], sub[1])
	}

	src = "#ABC#2#eee"
	sub, err = DouExstact(src, "#")
	if err != nil {
		t.Error("#ABC Exstact Error")
	}

	if len(sub) != 1 {
		t.Error("Exstact Error. Count Error")
	}

	if sub[0] != "ABC" {
		t.Errorf("Exstact Error. Value Error. 1:[%s]\n", sub[0])
	}
}

func TestExceptDouExstact(t *testing.T) {
	src := "#ABC"

	_, err := DouExstact(src, "#")
	if err == nil {
		t.Error("#ABC Exstact Error")
	}

}

func Benchmark_DouExstact(b *testing.B) {
	src := "#ABC#2#eee"
	b.StartTimer() //重新开始时间
	for i := 0; i < b.N; i++ {
		DouExstact(src, "#")
	}
}

func Benchmark_SymExstact(b *testing.B) {
	src := "{ABC}abc{DEF}jhi"
	b.StartTimer() //重新开始时间
	for i := 0; i < b.N; i++ {
		SymExstact(src, "{", "}")
	}
}

func ExampleSymExstact() {
	src := "{ABC}abc{DEF}jhi"
	sub, err := SymExstact(src, "{", "}")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(sub)
}

func ExampleDouExstact() {
	src := "#ABC#DEF#HHH#eee"

	sub, err := DouExstact(src, "#")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(sub)
}

func TestRemoveMultipeSpace(t *testing.T) {
	oldStr := "str1    str2      str3      str4     str5"
	str := RemoveMultipeSpace(oldStr)
	r := "str1 str2 str3 str4 str5"
	assert.Equal(t, r, str, "They should be equal")
	strs := strings.Split(str, " ")
	assert.Equal(t, 5, len(strs), "The length of strs should be 5")
}
