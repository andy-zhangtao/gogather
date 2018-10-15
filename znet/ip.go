package znet

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
)

// LocallIP 获取本地IP地址
func LocallIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network?")
}

// ConvertToHex 将IP地址转换为十六进制数据
// 例如将 127.0.0.1 转换为 7F 00 00 01
func ConvertToHex(ip string) (ipHex []byte, err error) {
	_ips := strings.Split(ip, ".")

	if len(_ips) != 4 {
		err = errors.New("Invalid IP")
		return
	}

	for _, i := range _ips {
		_i, err := strconv.Atoi(i)
		if err != nil {
			return ipHex, err
		}
		ipHex = append(ipHex, []byte(fmt.Sprintf("%02X", _i))...)
	}

	return
}
