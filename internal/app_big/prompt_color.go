package app_big

import (
	"fmt"
	"time"

	"github.com/go-color-term/go-color-term/coloring"
)

func Usage() {
	fmt.Println("Usage: qian [options] ...")
	fmt.Println("Options:")
	fmt.Println("  -v         \t 查看当前版本")
	fmt.Println("  -c         \t 新建PJSIP话机终端")
	fmt.Println("  -c [name]")
	fmt.Println("  -d num     \t 删除PJSIP话机")
	fmt.Println("  -p [page]  \t 提供书页, 运行对应页的案例")
	fmt.Println("  -h         \t 使用帮助")
	fmt.Println("")
}

func ShowVersion() {
	verNum := "1.0"
	year := time.Now().Year()
	ltd := "江苏仟略信息技术"

	ComName := coloring.Bold(coloring.BgMagenta(ltd))

	fmt.Println("laoqian " + coloring.Italic(verNum))
	fmt.Printf("Copyright (C) %d %s有限公司自制.\n", year, ComName)

	fmt.Println()
	fmt.Println("官方网站=> " + coloring.Blue("www.qianlue.cn"))
}

func ShowWarn() {

	print := coloring.New().White().Bold().Background().Yellow().Print()
	print("Use 'qian help' to learn how to use this tool :)\n")
}

func PrintWarn(text string) {
	print := coloring.New().White().Bold().Background().Yellow().Print()
	print(text)
}

func PrintError(text string) {
	print := coloring.New().White().Bold().Background().Red().Print()
	print(text)
	fmt.Println("Use 'qian help' to learn how to use this tool :)")
}
