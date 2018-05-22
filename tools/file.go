package tools

import (
	"bytes"
	"io"
)

//Write by zhangtao<ztao8607@gmail.com> . In 2018/3/30.

// LineCounter 统计文件行数
// r 通过os.Open获取的文件reader
func LineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}