package strings

import "testing"

func TestReverse(t *testing.T) {
	src := "www.mydomain.com.cn"
	dist := Reverse(src,".")

	if dist != "cn.com.mydomain.www" {
		t.Error("Conver Error "+dist)
	}

	src = "blog.www,domain.com.cn"
	dist = Reverse(src, ".")
	if dist != "cn.com.www,domain.blog" {
		t.Error("Conver Error "+dist)
	}
}

func TestReverseWithSeg(t *testing.T) {
	src := "www.mydomain.com.cn"
	dist := ReverseWithSeg(src,".","/")

	if dist != "cn/com/mydomain/www" {
		t.Error("Conver Error "+dist)
	}

	src = "blog.www,domain.com.cn"
	dist = ReverseWithSeg(src, ".","-")
	if dist != "cn-com-www,domain-blog" {
		t.Error("Conver Error "+dist)
	}
}