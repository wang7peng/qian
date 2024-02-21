package main

// This command will be more big, more full, and more colorful

import (
	"fmt"
	"os"
	Demo "qianDev/guidebook"
	L "qianDev/internal/app_big"
	F "qianDev/internal/app_small"
)

func init() {

}

func Usage() {
	fmt.Println("Usage: qian [options] ...")
	fmt.Println("Options:")
	fmt.Println("  -v         \t 查看当前版本")
	fmt.Println("  -c         \t 新建PJSIP话机终端")
	fmt.Println("  -c [name]")
	fmt.Println("  -d num     \t 删除PJSIP话机")
	fmt.Println("  -r         \t 控制话机沟通能力")
	fmt.Println("  -r '?---?' \t 同上, 通过形象串控制")
	fmt.Println("  -p [page]  \t 提供书页, 运行对应页的案例")
	fmt.Println("  -h         \t 使用帮助")
	fmt.Println("")
}

func ShowWarn() {
	fmt.Println("Use 'qian help' to learn how to use this tool :)")
}

func main() {
	if len(os.Args) == 1 {
		Usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "--version", "version", "-v", "-V":
		F.ShowVersion()

	case "--help", "help", "-h":
		Usage()

	case "--create", "-c":
		// 创建分机
		if len(os.Args) > 2 {
			F.CreateExtension(os.Args[2])
		} else {
			F.CreateEndPoint()
		}

	case "--delete", "-d":
		// 删除分机
		if len(os.Args) > 2 {
			F.DeleteExtension(os.Args[2])
		} else {
			ShowWarn()
		}

	case "--rule", "-r":
		// 控制话机之间的沟通关系
		if len(os.Args) == 2 {
			L.QA2ObeyRule()
			return
		}

		if len(os.Args) > 2 && len(os.Args[2]) > 7 {
			F.ObeyRule()
		} else {
			ShowWarn()
		}

	case "--page", "page", "-p", "-P":
		// relate page demo
		Demo.Run(os.Args[2])

	default:
		Usage()
	}
}
