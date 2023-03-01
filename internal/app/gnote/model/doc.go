package model

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Doc struct {
	Name          string   `json:"name"`
	FullPath      string   `json:"fullPath"`
	Contents      string   `json:"contents"`
	ContentsSplit []string `json:"contentsSplit"`
}

func (d *Doc) Query(fullPath string) (doc Doc, err error) {
	pwd, _ := os.Getwd()
	newPath := filepath.Join(pwd, "libs", "blog.liaosirui.com", fullPath)
	if !strings.HasSuffix(newPath, ".md") {
		newPath = newPath + ".md"
	}
	if !IsFile(newPath) {
		return doc, fmt.Errorf("文档 %v 不存在", newPath)
	}

	doc = Doc{}
	doc.ContentsSplit = make([]string, 0)

	// r, _ := os.Open(newPath)
	// defer r.Close()
	// s := bufio.NewScanner(r)
	// for s.Scan() {
	// 	doc.ContentsSplit = append(doc.ContentsSplit, s.Text())
	// }

	buf, _ := os.ReadFile(newPath)
	doc.Contents = string(buf)

	return
}

func IsFile(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !s.IsDir()
}
