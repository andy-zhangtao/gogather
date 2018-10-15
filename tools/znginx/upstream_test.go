package znginx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertUpstream(t *testing.T) {
	var nginx = `
server {
        server_name  tiny.cn;
        charset utf-8;
        ssl off;
        add_header Cache-Control no-store;
        error_page   500 502 503 504  /50x.html;


	location /logot {
        	proxy_pass http://1.tiny.cn;
	}
}
`
	var upstream = `
### [5b963b626f9d34000a41ed6e]-[/]-[upstream]-[start]
upstream 5b963b626f9d34000a41ed6e {
    server  192.168.1.178:8000;
}
### [5b963b626f9d34000a41ed6e]-[/]-[upstream]-[end]
`

	var result = `
### [5b963b626f9d34000a41ed6e]-[/]-[upstream]-[start]
upstream 5b963b626f9d34000a41ed6e {
    server  192.168.1.178:8000;
}
### [5b963b626f9d34000a41ed6e]-[/]-[upstream]-[end]

server {
        server_name  tiny.cn;
        charset utf-8;
        ssl off;
        add_header Cache-Control no-store;
        error_page   500 502 503 504  /50x.html;


	location /logot {
        	proxy_pass http://1.tiny.cn;
	}
}`
	newNginx := InsertUpstream(nginx, upstream)

	assert.Equal(t, result, newNginx)
}
