package app_small

import (
	"fmt"
	"time"
)

func Usage() {
	fmt.Println("Usage: qian [options] ...")
	fmt.Println("Options:")
	fmt.Println("  -v         \t 查看当前版本")
	fmt.Println("  -c num     \t 新建PJSIP话机")
	fmt.Println("  -d num     \t 删除PJSIP话机")
	fmt.Println("  -r '?'     \t 控制话机沟通能力")
	fmt.Println("             \t  A ----> B  \t A 能通 B 但 B 不一定通 A")
	fmt.Println("             \t  A ----- B  \t A B 能互打")
	fmt.Println("             \t  A --x-- B  \t A B 不能打")
	fmt.Println("  -h         \t 使用帮助")
	fmt.Println("")
}

func ShowWarn() {
	fmt.Println("Use 'qian help' to learn how to use this tool :)")
}

func ShowVersion() {
	year := time.Now().Year()
	ltd := "江苏仟略信息技术"

	fmt.Println("qian 1.0")
	fmt.Printf("Copyright (C) %d %s有限公司自制.\n", year, ltd)

	fmt.Println()
	fmt.Println("Mail bug or suggestions to <www.qianlue.cn>")
}
