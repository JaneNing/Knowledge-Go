package main

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func main() {
	fmt.Println(PrintLastDirName(PrintFileAbsolutePath()))
}

// 输出当前文件绝对路径
func PrintFileAbsolutePath() string {
	_, filePath, _, _ := runtime.Caller(0)
	return filePath
}

// 输出输入地址的绝对地址
func PrintDirAbsolutePath(path string) string {
	if path == "" {
		path = "."
	}
	filePath, _ := filepath.Abs(path)
	return filePath
}

// 输出输入路径的上一级目录
func PrintAbsolutePath(path string) string {
	if path == "" {
		path = "."
	}
	return filepath.Dir(PrintDirAbsolutePath(path))
}

// 输出输入路径的的上一级目录的最后一个文件夹名
func PrintLastDirName(path string) string {
	if path == "" {
		path = "."
	}
	return filepath.Base(PrintAbsolutePath(path))
}
