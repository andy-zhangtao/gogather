FROM golang:1.10-alpine3.7
LABEL MAINTAINER=ztao@gmail.com
RUN		apk update && \
		apk add git expect curl && \
		go get -u github.com/golang/dep/cmd/dep
RUN mkdir -p /go/src/github.com/andy-zhangtao/gogather
COPY unit-test.sh /go/src/github.com/andy-zhangtao/gogather/unit-test.sh
ADD convert  /go/src/github.com/andy-zhangtao/gogather/convert
ADD crypto /go/src/github.com/andy-zhangtao/gogather/crypto
ADD random /go/src/github.com/andy-zhangtao/gogather/random
ADD strings /go/src/github.com/andy-zhangtao/gogather/strings
ADD time /go/src/github.com/andy-zhangtao/gogather/time
ADD tools /go/src/github.com/andy-zhangtao/gogather/tools
ADD zReflect /go/src/github.com/andy-zhangtao/gogather/zReflect
ADD zlog /go/src/github.com/andy-zhangtao/gogather/zlog
ADD znet /go/src/github.com/andy-zhangtao/gogather/znet
ADD zoutput /go/src/github.com/andy-zhangtao/gogather/zoutput
ADD zsort /go/src/github.com/andy-zhangtao/gogather/zsort

WORKDIR /go/src/github.com/andy-zhangtao/gogather
ENTRYPOINT ["sh","unit-test.sh"]
