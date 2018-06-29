package znginx

import (
	"regexp"
	"strings"
	"bufio"
	"fmt"
)

//ExtractHosts 从nginx server片段中抽取domain
//nginx server片段
func ExtractHosts(nginx string) (hosts []string) {
	re := regexp.MustCompile("server_name(.*);")
	h := re.FindStringSubmatch(nginx)
	if len(h) == 1 {
		return
	}

	_hosts := strings.Split(strings.TrimSpace(h[1]), " ")
	for _, h := range _hosts {
		if strings.TrimSpace(h) != "" {
			hosts = append(hosts, h)
		}
	}
	return
}

//ExtractUpstream 从nginx server片段中抽取upstream片段
//nginx server片段
func ExtractUpstream(nginx string) (upstream []string) {
	scanner := bufio.NewScanner(strings.NewReader(nginx))

	hasUps := false
	var ups string = ""

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "upstream") {
			hasUps = true
			ups = scanner.Text() + "\n"
			continue
		}

		if hasUps {
			ups += scanner.Text() + "\n"
			if strings.Contains(scanner.Text(), "}") {
				hasUps = false
				upstream = append(upstream, ups)
			}
		}
	}

	return
}

//ExtractLocation 从nginx server片段中抽取domain和location的映射数据
func ExtractLocation(nginx string) (location map[string][]string) {
	hosts := ExtractHosts(nginx)

	location = make(map[string][]string)

	loc := extractLocation(nginx)

	for _, h := range hosts {
		location[h] = loc
	}

	return
}

func extractLocation(nginx string) (location []string) {
	scanner := bufio.NewScanner(strings.NewReader(nginx))

	hasLoc := false
	var loc string = ""

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "location") {
			hasLoc = true
			loc = scanner.Text() + "\n"
			continue
		}

		if hasLoc {
			loc += scanner.Text() + "\n"
			if strings.Contains(scanner.Text(), "}") {
				hasLoc = false
				location = append(location, loc)
			}
		}
	}

	return
}

//MergeServerF1 标准Server合并模式
//每个server片段中server_name 只能是单域名，在多域名情况下使用此函数会导致合并后的业务语义错误
//通过isMerge返回是否发生了合并行为
func MergeServerF1(nginx1, nginx2 string, nginxs ...string) (nginx string, isMerge bool, err error) {
	isMerge = false

	host1 := ExtractHosts(nginx1)
	ups2 := ExtractUpstream(nginx2)
	host2 := ExtractHosts(nginx2)

	if len(host1) > 1 || len(host2) > 1 {
		fmt.Sprintf("MergeServerF1 Only Merge Single Domain Server [%s] [%s] ", strings.Join(host1, " "), strings.Join(host2, " "))
		return
	}

	if host1[0] == host2[0] {
		isMerge = true
		locMap2 := ExtractLocation(nginx2)
		// 需要合并

		// 先合并location
		nginx1 = strings.TrimSpace(nginx1)

		_nginx := nginx1[:strings.LastIndex(nginx1, "}")]
		_nginx += "\n"
		for _, l := range locMap2[host2[0]] {
			_nginx += l
		}
		_nginx += "}\n"

		// 再合并upstream
		nginx = strings.Join(ups2, "\n") + _nginx

	}

	for _, n := range nginxs {
		host := ExtractHosts(n)
		if len(host) > 1 {
			fmt.Sprintf("MergeServerF1 Only Merge Single Domain Server [%s]", strings.Join(host, " "))
			continue
		}

		if host1[0] == host[0] {
			isMerge = true
			locMap := ExtractLocation(n)
			ups := ExtractUpstream(n)

			_nginx := nginx[:strings.LastIndex(nginx, "}")]

			_nginx += "\n"

			for _, l := range locMap[host[0]] {
				_nginx += l
			}

			_nginx += "}\n"

			nginx = strings.Join(ups, "\n") + _nginx
		}

	}

	nginx += "\n#Auto Merge By GoGather \n"
	return
}
