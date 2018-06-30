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
