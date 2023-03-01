package model

import (
	"errors"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type Path struct {
	Name     string   `json:"name"`
	FullPath string   `json:"fullPath"`
	SubDocs  []string `json:"subDocs"`
	SubPaths []string `json:"subPaths"`
}

func (p *Path) Query(fullPath string) (path Path, err error) {
	pwd, _ := os.Getwd()
	newPath := filepath.Join(pwd, "libs", "blog.liaosirui.com", fullPath)
	if !IsDir(newPath) {
		return path, errors.New("文件夹不存在")
	}
	path = Path{
		FullPath: strings.Clone(fullPath),
		Name:     filepath.Base(fullPath),
	}

	files, _ := os.ReadDir(newPath)
	path.SubPaths = make([]string, 0)
	path.SubDocs = make([]string, 0)

	for _, f := range files {
		if f.IsDir() {
			dirName := strings.Clone(f.Name())
			if dirName == ".assets" {
				continue
			}
			path.SubPaths = append(path.SubPaths, dirName)
			continue
		}
		fileName := strings.Clone(f.Name())
		if filepath.Ext(fileName) != ".md" {
			continue
		}
		fileName = strings.Replace(fileName, ".md", "", -1)
		if fileName == "README" {
			continue
		}
		path.SubDocs = append(path.SubDocs, fileName)
	}

	sort.Strings(path.SubPaths)
	sort.Strings(path.SubDocs)

	return
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}
