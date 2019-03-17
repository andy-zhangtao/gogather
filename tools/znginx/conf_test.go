package znginx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractHosts(t *testing.T) {
	var nginx = `
upstream tomcats9540 {
    server 192.168.1.216:8000;
}

#admin
server {
	server_name  www.chinazt.cc  blog.chinazt.cc;

	location / {
		index index.html index.htm index.jsp;
		proxy_pass         http://tomcats9540;
		proxy_redirect     off;
		proxy_set_header   Host             $host;
		proxy_set_header   X-Real-IP        $remote_addr;
		proxy_set_header  X-Forwarded-For  $proxy_add_x_forwarded_for;
		client_max_body_size       10m;
		client_body_buffer_size    128k;
		proxy_connect_timeout      90;
		proxy_send_timeout         90;
		proxy_read_timeout         90;
		proxy_buffer_size          4k;
		proxy_buffers              4 32k;
		proxy_busy_buffers_size    64k;
		proxy_temp_file_write_size 64k;
	}
}

`

	hosts := ExtractHosts(nginx)

	assert.Equal(t, 2, len(hosts))
	assert.EqualValues(t, "www.chinazt.cc", hosts[0])
	assert.EqualValues(t, "blog.chinazt.cc", hosts[1])
}

func TestExtraceUpstreamValue(t *testing.T) {
	var nginx = `
	upstream tomcats9540 {
		server 192.168.1.216:8000;
		server 192.168.1.216:9000;
		# server 192.168.1.216:9000;
	}


	#admin
	server {
		server_name  www.chinazt.cc  blog.chinazt.cc;

		location / {
			index index.html index.htm index.jsp;
			proxy_pass         http://tomcats9540;
			proxy_redirect     off;
			proxy_set_header   Host             $host;
			proxy_set_header   X-Real-IP        $remote_addr;
			proxy_set_header  X-Forwarded-For  $proxy_add_x_forwarded_for;
			client_max_body_size       10m;
			client_body_buffer_size    128k;
			proxy_connect_timeout      90;
			proxy_send_timeout         90;
			proxy_read_timeout         90;
			proxy_buffer_size          4k;
			proxy_buffers              4 32k;
			proxy_busy_buffers_size    64k;
			proxy_temp_file_write_size 64k;
		}
	}
	`
	ups := ExtractUpstream(nginx)
	server, err := ExtractUpstreamValue(ups[0])
	assert.Nil(t, err)
	assert.Equal(t, 2, len(server))
	assert.Equal(t, "192.168.1.216:8000", server[0])
	assert.Equal(t, "192.168.1.216:9000", server[1])
}
func TestExtractUpstream1(t *testing.T) {
	var nginx = `### [5bdfe6df67609d000a5e5c4b]-[/]-[upstream]-[start]
	upstream 5bdfe6df67609d000a5e5c4b {
		server  192.168.1.237:8000;
	}
	### [5bdfe6df67609d000a5e5c4b]-[/]-[upstream]-[end]

	server{

		server_name qa.tiny.eqshow.cn;




		location / {
			proxy_pass http://5bdfe6df67609d000a5e5c4b;
			proxy_set_header X-Real-IP $remote_addr;
			proxy_set_header Host $host;
			proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;



		}

	}`

	ups := ExtractUpstream(nginx)
	assert.Equal(t, 1, len(ups))
}
func TestExtractUpstream(t *testing.T) {
	var nginx = `
upstream tomcats9540 {
    server 192.168.1.216:8000;
}


#admin
server {
	server_name  www.chinazt.cc  blog.chinazt.cc;

	location / {
		index index.html index.htm index.jsp;
		proxy_pass         http://tomcats9540;
		proxy_redirect     off;
		proxy_set_header   Host             $host;
		proxy_set_header   X-Real-IP        $remote_addr;
		proxy_set_header  X-Forwarded-For  $proxy_add_x_forwarded_for;
		client_max_body_size       10m;
		client_body_buffer_size    128k;
		proxy_connect_timeout      90;
		proxy_send_timeout         90;
		proxy_read_timeout         90;
		proxy_buffer_size          4k;
		proxy_buffers              4 32k;
		proxy_busy_buffers_size    64k;
		proxy_temp_file_write_size 64k;
	}
}

upstream tomcats9541 {
    server 192.168.1.216:8000;
}

`

	ups := ExtractUpstream(nginx)

	assert.Equal(t, 2, len(ups))
}

func TestExtractLocation1(t *testing.T) {
	var nginx = `http {
		include	/home/eqs/soft/nginx/conf/mime.types;
		default_type  application/octet-stream;
		server_tokens off;

		include common-log.conf;

		sendfile	on;
		#tcp_nopush     on;

		keepalive_timeout  65;

		client_max_body_size 32M;

		proxy_buffering		on;
		proxy_buffer_size	4k;
		proxy_buffers		512 4k;
		proxy_busy_buffers_size	64k;
		proxy_cache_path	/data/eqs/proxy_cache levels=1:2 keys_zone=cache:300m inactive=24h max_size=10g use_temp_path=off;

		gzip	on;
		gzip_min_length	1k;
		gzip_buffers	4 16k;
		gzip_comp_level	2;
		gzip_types	text/plain text/css text/javascript application/json application/javascript application/x-javascript application/xml;
		gzip_vary	off;
		upstream ups-caas {
			server 192.168.0.18:80 max_fails=2 fail_timeout=5s weight=10;
		}
		include conf.d/*.conf;
	}`

	location := ExtractLocation(nginx)
	assert.Equal(t, 0, len(location))
}
func TestExtractLocation(t *testing.T) {
	var nginx = `
upstream tomcats9540 {
    server 192.168.1.216:8000;
}


#admin
server {
	server_name  www.chinazt.cc  blog.chinazt.cc;

	location / {
		index index.html index.htm index.jsp;
		proxy_pass         http://tomcats9540;
		proxy_redirect     off;
		proxy_set_header   Host             $host;
		proxy_set_header   X-Real-IP        $remote_addr;
		proxy_set_header  X-Forwarded-For  $proxy_add_x_forwarded_for;
		client_max_body_size       10m;
		client_body_buffer_size    128k;
		proxy_connect_timeout      90;
		proxy_send_timeout         90;
		proxy_read_timeout         90;
		proxy_buffer_size          4k;
		proxy_buffers              4 32k;
		proxy_busy_buffers_size    64k;
		proxy_temp_file_write_size 64k;
	}

	location /login {
		index index.html index.htm index.jsp;
		proxy_pass         http://tomcats9540;
		proxy_redirect     off;
		proxy_set_header   Host             $host;
		proxy_set_header   X-Real-IP        $remote_addr;
		proxy_set_header  X-Forwarded-For  $proxy_add_x_forwarded_for;
		client_max_body_size       10m;
		client_body_buffer_size    128k;
		proxy_connect_timeout      90;
		proxy_send_timeout         90;
		proxy_read_timeout         90;
		proxy_buffer_size          4k;
		proxy_buffers              4 32k;
		proxy_busy_buffers_size    64k;
		proxy_temp_file_write_size 64k;
	}
}

`

	location := ExtractLocation(nginx)

	assert.EqualValues(t, 2, len(location))
	assert.EqualValues(t, 2, len(location["www.chinazt.cc"]))
	assert.EqualValues(t, 2, len(location["blog.chinazt.cc"]))
}

func TestMergeServerF1(t *testing.T) {
	var nginx1 = `
upstream tomcats9540 {
    server 192.168.1.216:8000;
}


#admin
server {
	server_name  www.chinazt.cc;

	location / {
		index index.html index.htm index.jsp;
		proxy_pass         http://tomcats9540;
		proxy_redirect     off;
		proxy_set_header   Host             $host;
		proxy_set_header   X-Real-IP        $remote_addr;
		proxy_set_header  X-Forwarded-For  $proxy_add_x_forwarded_for;
		client_max_body_size       10m;
		client_body_buffer_size    128k;
		proxy_connect_timeout      90;
		proxy_send_timeout         90;
		proxy_read_timeout         90;
		proxy_buffer_size          4k;
		proxy_buffers              4 32k;
		proxy_busy_buffers_size    64k;
		proxy_temp_file_write_size 64k;
	}
}
`
	var nginx2 = `
upstream tomcats9541 {
    server 192.168.1.213:8000;
}


#admin
server {
	server_name  www.chinazt.cc;

	location ~* /login {
		index index.html index.htm index.jsp;
		proxy_pass         http://tomcats9541;
		proxy_redirect     off;
		proxy_set_header   Host             $host;
		proxy_set_header   X-Real-IP        $remote_addr;
		proxy_set_header  X-Forwarded-For  $proxy_add_x_forwarded_for;
		client_max_body_size       10m;
		client_body_buffer_size    128k;
		proxy_connect_timeout      90;
		proxy_send_timeout         90;
		proxy_read_timeout         90;
		proxy_buffer_size          4k;
		proxy_buffers              4 32k;
		proxy_busy_buffers_size    64k;
		proxy_temp_file_write_size 64k;
	}
}
`
	var nginx3 = `
	upstream tomcats9542 {
	   server 192.168.1.211:8000;
	}


	#admin
	server {
		server_name  www.chinazt.cc;

		location ~* /logout {
			index index.html index.htm index.jsp;
			proxy_pass         http://tomcats9542;
			proxy_redirect     off;
			proxy_set_header   Host             $host;
			proxy_set_header   X-Real-IP        $remote_addr;
			proxy_set_header  X-Forwarded-For  $proxy_add_x_forwarded_for;
			client_max_body_size       10m;
			client_body_buffer_size    128k;
			proxy_connect_timeout      90;
			proxy_send_timeout         90;
			proxy_read_timeout         90;
			proxy_buffer_size          4k;
			proxy_buffers              4 32k;
			proxy_busy_buffers_size    64k;
			proxy_temp_file_write_size 64k;
		}
	}
	`

	nginx, isMerge, err := MergeServerF1(nginx1, nginx2)
	assert.Nil(t, err)
	locMap := ExtractLocation(nginx)
	assert.EqualValues(t, 1, len(locMap))
	assert.EqualValues(t, true, isMerge)
	assert.EqualValues(t, 2, len(locMap["www.chinazt.cc"]))

	nginx, _, err = MergeServerF1(nginx1, nginx2, nginx3)
	assert.Nil(t, err)
	locMap = ExtractLocation(nginx)
	assert.EqualValues(t, 1, len(locMap))
	assert.EqualValues(t, true, isMerge)
	assert.EqualValues(t, 3, len(locMap["www.chinazt.cc"]))

}

func TestExtractLocationDestProxyPass(t *testing.T) {
	var nginx = `
	upstream tomcats9542 {
	   server 192.168.1.211:8000;
	}


	#admin
	server {
		server_name  www.chinazt.cc;

		location ~* /logout {
			index index.html index.htm index.jsp;
			proxy_pass         http://tomcats9542;
			proxy_redirect     off;
			proxy_set_header   Host             $host;
			proxy_set_header   X-Real-IP        $remote_addr;
			proxy_set_header  X-Forwarded-For  $proxy_add_x_forwarded_for;
			client_max_body_size       10m;
			client_body_buffer_size    128k;
			proxy_connect_timeout      90;
			proxy_send_timeout         90;
			proxy_read_timeout         90;
			proxy_buffer_size          4k;
			proxy_buffers              4 32k;
			proxy_busy_buffers_size    64k;
			proxy_temp_file_write_size 64k;
		}
	}
	`

	location := ExtractLocation(nginx)
	assert.Equal(t, 1, len(location))

	dest, isroot, loc := ExtractLocationDest(location["www.chinazt.cc"][0])

	assert.Equal(t, false, isroot)
	assert.Equal(t, "~* /logout", loc)
	assert.Equal(t, "http://tomcats9542", dest)
}

func TestExtractLocationDestRoot(t *testing.T) {
	var nginx = `
	upstream tomcats9542 {
	   server 192.168.1.211:8000;
	}


	#admin
	server {
		server_name  www.chinazt.cc;

		location ~* /logout {
			index index.html index.htm index.jsp;
			root         /var/www/root;
			proxy_redirect     off;
			proxy_set_header   Host             $host;
			proxy_set_header   X-Real-IP        $remote_addr;
			proxy_set_header  X-Forwarded-For  $proxy_add_x_forwarded_for;
			client_max_body_size       10m;
			client_body_buffer_size    128k;
			proxy_connect_timeout      90;
			proxy_send_timeout         90;
			proxy_read_timeout         90;
			proxy_buffer_size          4k;
			proxy_buffers              4 32k;
			proxy_busy_buffers_size    64k;
			proxy_temp_file_write_size 64k;
		}
	}
	`

	location := ExtractLocation(nginx)
	assert.Equal(t, 1, len(location))

	dest, isroot, loc := ExtractLocationDest(location["www.chinazt.cc"][0])

	assert.Equal(t, true, isroot)
	assert.Equal(t, "~* /logout", loc)
	assert.Equal(t, "/var/www/root", dest)
}

func TestExtractByComment(t *testing.T) {
	var nginx = `
### [ups] start ###
upstream tomcats9540 {
    server 192.168.1.216:8000;
}
### [ups] end ###

#admin
server {
	server_name  www.chinazt.cc;

	### [location] start ###
	location / {
		index index.html index.htm index.jsp;
		proxy_pass         http://tomcats9540;
		proxy_redirect     off;
		proxy_set_header   Host             $host;
		proxy_set_header   X-Real-IP        $remote_addr;
		proxy_set_header  X-Forwarded-For  $proxy_add_x_forwarded_for;
		client_max_body_size       10m;
		client_body_buffer_size    128k;
		proxy_connect_timeout      90;
		proxy_send_timeout         90;
		proxy_read_timeout         90;
		proxy_buffer_size          4k;
		proxy_buffers              4 32k;
		proxy_busy_buffers_size    64k;
		proxy_temp_file_write_size 64k;
	}
	### [location] end ###
}
`

	content, err := ExtractByComment(nginx, "### [ups] start ###", "### [ups] end ###")
	assert.Nil(t, err)
	assert.EqualValues(t, 3, len(content))

	content, err = ExtractByComment(nginx, "### [ups] start ###")
	assert.Error(t, err)
	assert.EqualValues(t, 0, len(content))

	content, err = ExtractByComment(nginx, "### [ups] start ###", "### [ups] end ###", "### [location] start ###", "### [location] end ###")
	assert.Nil(t, err)
	assert.EqualValues(t, 20, len(content))
}

func TestExtractAndReplaceByComment(t *testing.T) {
	var nginx = `
### [ups] start ###
upstream tomcats9540 {
    server 192.168.1.216:8000;
}
### [ups] end ###

#admin
server {
	server_name  www.chinazt.cc;

	### [location] start ###
	location / {
		index index.html index.htm index.jsp;
		proxy_pass         http://tomcats9540;
		proxy_redirect     off;
		proxy_set_header   Host             $host;
		proxy_set_header   X-Real-IP        $remote_addr;
		proxy_set_header  X-Forwarded-For  $proxy_add_x_forwarded_for;
		client_max_body_size       10m;
		client_body_buffer_size    128k;
		proxy_connect_timeout      90;
		proxy_send_timeout         90;
		proxy_read_timeout         90;
		proxy_buffer_size          4k;
		proxy_buffers              4 32k;
		proxy_busy_buffers_size    64k;
		proxy_temp_file_write_size 64k;
	}
	### [location] end ###
}
`
	replace := []string{
		"upstream tomcats9540 {",
		"server 192.168.1.222:8000;",
		"}",
	}

	isReplace, _, err := ExtractAndReplaceByComment(nginx, replace, "### [ups] start ###", "### [ups] end ###")
	assert.Nil(t, err)
	assert.EqualValues(t, true, isReplace)
}
