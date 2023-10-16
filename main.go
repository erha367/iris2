package main

import (
	"flag"
	"iris/lib"
	"strings"
)

var url, color string

func main() {
	//参数解析
	flag.StringVar(&url, "url", "", "输入图片的绝对路径")
	flag.StringVar(&color, "color", "default", "颜色：red,blue,green,yellow")
	flag.Parse()
	//解析
	urlArr := strings.Split(url, `.`)
	tail := urlArr[len(urlArr)-1]
	if tail == "png" {
		lib.DrawPng(url, color)
	} else {
		lib.DrawJpg(url, color)
	}

	return
}
