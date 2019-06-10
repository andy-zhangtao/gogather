package znet

import (
	"fmt"
	"strconv"
)

// ConvertIP 将网络字节序的IP地址转换成人可识别的IP地址
func ConvertIP(iphex string) (ip string, err error) {

	for i := 0; i <= len(iphex)-2; i += 2 {
		p, err := strconv.ParseInt(iphex[i:i+2], 16, 10)
		if err != nil {
			return ip, err
		}

		ip = fmt.Sprintf(".%d%s", p, ip)
	}

	return ip[1:], nil
}
