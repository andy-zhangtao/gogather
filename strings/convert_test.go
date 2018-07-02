package strings

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	src := "www.mydomain.com.cn"
	dist := Reverse(src, ".")

	if dist != "cn.com.mydomain.www" {
		t.Error("Conver Error " + dist)
	}

	src = "blog.www,domain.com.cn"
	dist = Reverse(src, ".")
	if dist != "cn.com.www,domain.blog" {
		t.Error("Conver Error " + dist)
	}
}

func TestReverseWithSeg(t *testing.T) {
	src := "www.mydomain.com.cn"
	dist := ReverseWithSeg(src, ".", "/")

	if dist != "cn/com/mydomain/www" {
		t.Error("Conver Error " + dist)
	}

	src = "blog.www,domain.com.cn"
	dist = ReverseWithSeg(src, ".", "-")
	if dist != "cn-com-www,domain-blog" {
		t.Error("Conver Error " + dist)
	}
}

func TestReplaceAscii(t *testing.T) {
	var src = `\n### [tiny]-[/logout]-[upstream]-[start]\nupstream 1.tiny.cn {\n    server  192.168.1.52:8000;\n}\n### [tiny]-[/logout]-[upstream]-[end]\n\n### [tiny-login]-[/login]-[upstream]-[start]\nupstream 2.tiny.cn {\n    server  192.168.1.129:8000;\n}\n### [tiny-login]-[/login]-[upstream]-[end]\n\n### [tiny-root]-[/]-[upstream]-[start]\nupstream 2.tiny.cn {\n    server  192.168.1.129:8000;\n}\n### [tiny-root]-[/]-[upstream]-[end]\n\nserver {\n        server_name  tiny.cn;\n        charset utf-8;\n        ssl off;\n        add_header Cache-Control no-store;\n        error_page 404 = http://mp.com/#/error/404;\n        error_page   500 502 503 504  /50x.html;\n\n\n\tlocation /logot {\n        \tproxy_pass http://1.tiny.cn;\n\t}\n\n\tlocation /login {\n\n\t\t\tproxy_pass http://2.tiny.cn;\t\n\t}\n\n}\n`
	src = ReplaceAscii(src, []string{"\\n", "\\t"})

	var expect = `
### [tiny]-[/logout]-[upstream]-[start]
upstream 1.tiny.cn {
    server  192.168.1.52:8000;
}
### [tiny]-[/logout]-[upstream]-[end]

### [tiny-login]-[/login]-[upstream]-[start]
upstream 2.tiny.cn {
    server  192.168.1.129:8000;
}
### [tiny-login]-[/login]-[upstream]-[end]

### [tiny-root]-[/]-[upstream]-[start]
upstream 2.tiny.cn {
    server  192.168.1.129:8000;
}
### [tiny-root]-[/]-[upstream]-[end]

server {
        server_name  tiny.cn;
        charset utf-8;
        ssl off;
        add_header Cache-Control no-store;
        error_page 404 = http://mp.com/#/error/404;
        error_page   500 502 503 504  /50x.html;


	location /logot {
        	proxy_pass http://1.tiny.cn;
	}

	location /login {

			proxy_pass http://2.tiny.cn;	
	}

}
`
	assert.Equal(t, expect, src)
}
