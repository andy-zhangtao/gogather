package znginx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertLocation(t *testing.T) {
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
	var login = `
	location /login{
		proxy_pass http://login.cn;
	}
`

	var result = `server {
        server_name  tiny.cn;
        charset utf-8;
        ssl off;
        add_header Cache-Control no-store;
        error_page   500 502 503 504  /50x.html;


	location /logot {
        	proxy_pass http://1.tiny.cn;
	}



	location /login{
		proxy_pass http://login.cn;
	}

}`
	newNginx := InsertLocation(nginx, login)

	assert.Equal(t, result, newNginx)
}
