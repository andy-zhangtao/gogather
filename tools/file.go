package tools

import (
	"bytes"
	"io"
	"fmt"
	"os"
	"io/ioutil"
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

// CopyDir copies a dir from src to dst.
// src should be a full path. Also dst too.
// If src is a file, then will invoke CopyFile.
// If src is a dir, then will copy all the files it contains to dst.
/**
##### Example

```go
package main

import (
	"github.com/andy-zhangtao/gogather/tools"
	"fmt"
)

func main() {
	err := tools.CopyDir("/Users/zhangtao/SourceCode/golang/go/src/temp/test", "/tmp/test")
	if err != nil {
		fmt.Println(err)
	}
}

```
*/
func CopyDir(src, dst string) (err error) {

	fi, err := os.Stat(src)
	if err != nil {
		return
	}

	switch mode := fi.Mode(); {
	case mode.IsDir():
		return copyDir(src, dst)
	case mode.IsRegular():
		return CopyFile(src, dst)
	}
	return
}

func copyDir(src, dst string) (err error) {
	if _, err := os.Stat(dst); os.IsNotExist(err) {
		os.Mkdir(dst, 0777)
	}

	files, err := ioutil.ReadDir(src)
	if err != nil {
		return
	}

	for _, f := range files {
		var err error
		switch mode := f.Mode(); {
		case mode.IsDir():
			err = copyDir(fmt.Sprintf("%s/%s", src, f.Name()), fmt.Sprintf("%s/%s", dst, f.Name()))
		case mode.IsRegular():
			err = CopyFile(fmt.Sprintf("%s/%s", src, f.Name()), fmt.Sprintf("%s/%s", dst, f.Name()))
		}

		if err != nil {
			return err
		}
	}

	return
}

// CopyFile copies a file from src to dst. If src and dst files exist, and are
// the same, then return success. Otherise, attempt to create a hard link
// between the two files. If that fail, copy the file contents from src to dst.
func CopyFile(src, dst string) (err error) {
	sfi, err := os.Stat(src)
	if err != nil {
		return
	}
	if !sfi.Mode().IsRegular() {
		// cannot copy non-regular files (e.g., directories,
		// symlinks, devices, etc.)
		return fmt.Errorf("CopyFile: non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
	}
	dfi, err := os.Stat(dst)
	if err != nil {
		if !os.IsNotExist(err) {
			return
		}
	} else {
		if !(dfi.Mode().IsRegular()) {
			return fmt.Errorf("CopyFile: non-regular destination file %s (%q)", dfi.Name(), dfi.Mode().String())
		}
		if os.SameFile(sfi, dfi) {
			return
		}
	}
	if err = os.Link(src, dst); err == nil {
		return
	}
	err = copyFileContents(src, dst)
	return
}

// copyFileContents copies the contents of the file named src to the file named
// by dst. The file will be created if it does not already exist. If the
// destination file exists, all it's contents will be replaced by the contents
// of the source file.
func copyFileContents(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}
