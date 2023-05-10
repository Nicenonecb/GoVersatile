package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
	"strings"
)

func printDirectoryStructure(root string, printFunc func(string)) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("遇到错误: %v\n", err)
			return err
		}

		// 计算缩进
		indent := strings.Repeat("│   ", strings.Count(path, string(os.PathSeparator))-1)
		if info.IsDir() {
			// 打印目录
			indent = strings.TrimSuffix(indent, "│   ")
			indent += "├── "
		} else {
			// 打印文件
			indent += "└── "
		}

		printFunc(fmt.Sprintf("%s%s", indent, filepath.Base(path)))
		return nil
	})
}

func main() {
	root := "/Users/justin/Desktop/GoVersatile" // 你可以将这个值替换为你想要遍历的目录

	err := printDirectoryStructure(root, func(s string) {
		fmt.Println(s)
	})

	if err != nil {
		fmt.Printf("filepath.Walk 遇到错误: %v\n", err)
	}
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Gi!",
		})
	})
	r.Run(":9090")
}
