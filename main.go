package main

import (
	_ "base-learn/matchers" //初始化rss的match
	"base-learn/search"
	"log"
	"os"
)

// init()加载文件之后调用
func init() {
	// 将日志输出到标准输出
	log.SetOutput(os.Stdout)
}

// main是整个程序的入口
func main() {
	// 使用特定的项做搜索
	search.Run("President")
	search.Run("Presidential")
}
