package znginx

import (
	"regexp"
	"strings"
	"bufio"
	"fmt"
	"errors"
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

//ExtractByComment 提取指定注释间的数据
//comment必须成对出现。即 comment-start, comment-end, comment-start, comment-end...
//具体使用方式可参考conf_test.go中的实例代码
func ExtractByComment(nginx string, comment ...string) (content []string, err error) {
	if len(comment)%2 != 0 {
		err = errors.New("comment must be pairing. start, end, start, end ...")
		return
	}

	for i := 0; i < len(comment); i = i + 2 {
		_, _c := extractByComment(nginx, comment[i], comment[i+1])
		content = append(content, _c ...)
	}

	return
}

//ExtractAndReplaceByComment 替换指定注释之间的数据, 仅支持替换一个注释区间的数据
//replace 准备替换的文件内容
func ExtractAndReplaceByComment(nginx string, replace []string, comment ...string, ) (isReplace bool, newNginx string, err error) {
	if len(comment)%2 != 0 {
		err = errors.New("comment must be pairing. start, end, start, end ...")
		return
	}

	var line []int
	for i := 0; i < len(comment); i = i + 2 {
		_l, _ := extractByComment(nginx, comment[i], comment[i+1])
		line = append(line, _l...)
	}

	newNginx = replaceNginxContent(nginx, line, replace)

	return !(nginx == newNginx), newNginx, nil
}

// extractByComment 提取指定注释之间的数据
// nginx 配置数据内容
// comment1 注释起始内容
// comment2 注释结束内容
// 如果提取到内容，则返回内容在nginx数据中的行数，同时返回所提取到的数据。 例如:
// nginx:
// 15. ### [comment-start] ###
// 16. upstream {
// 17. 	server 127.0.0.1;
// 18. }
// 19. ### [comment-end] ###
// 则返回line=[16,17,18]. content=["upstream {","server 127.0.0.1;","}"]
// 返回的content保留原始格式
func extractByComment(nginx, comment1, comment2 string) (line []int, content []string) {
	scanner := bufio.NewScanner(strings.NewReader(nginx))

	findContent := false
	idx := 0
	for scanner.Scan() {
		idx++
		if strings.HasPrefix(strings.TrimSpace(scanner.Text()), "#") && strings.Contains(scanner.Text(), comment1) {
			findContent = true
			continue
		}

		if strings.HasPrefix(strings.TrimSpace(scanner.Text()), "#") && strings.Contains(scanner.Text(), comment2) {
			return
		}

		if findContent {
			line = append(line, idx)
			content = append(content, scanner.Text())
		}

	}

	return
}

// replaceNginxContent 替换Nginx指定行号的内容. 此函数会将content内容插入在line[0]的位置上
// nginx 数据源
// line 准备要修改的行号
// content 替换数据源
// 操作时会将指定行号的内容删除，然后替换content的数据
// 因此content的数据量不必和行号相等
// 最后返回新的nginx数据.
func replaceNginxContent(nginx string, line []int, content []string) (newNginx string) {
	_lineMap := make(map[int]int)
	for _, l := range line {
		_lineMap[l] = l
	}

	scanner := bufio.NewScanner(strings.NewReader(nginx))

	hasReplace := false
	idx := 0
	for scanner.Scan() {
		idx++

		if _, ok := _lineMap[idx]; ok {
			if !hasReplace {
				for _, c := range content {
					newNginx += c + "\n"
				}
				hasReplace = true
			}
		} else {
			newNginx += scanner.Text() + "\n"
		}
	}

	return newNginx
}
