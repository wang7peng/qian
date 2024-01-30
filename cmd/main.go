package main

import (
	"fmt"
	"os"
	Demo "qianDev/guidebook"
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
	fmt.Println("  -p [page]  \t 提供书页, 运行对应页的案例")
	fmt.Println("  -h         \t 使用帮助")
	fmt.Println("")
	fmt.Println("Use 'qian help' to learn how to use this tool :)")
}

func main() {
	if len(os.Args) == 1 {
		Usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "--version", "version", "-v", "-V":
		fmt.Println("qian 1.0")
		fmt.Println("Copyright (C) 2024 江苏仟略有限公司自制.")

	case "--help", "help", "-h":
		Usage()

	case "--create", "-c":
		// 创建分机
		if len(os.Args) > 2 {
			CreateExtension(os.Args[2])
		} else {
			CreateEndPoint()
		}

	case "--delete", "-d":
		// 删除分机
		if len(os.Args) > 2 {
			DeleteExtension(os.Args[2])
		} else {
			Usage()
		}

	case "--page", "page", "-p", "-P":
		// relate page demo
		Demo.Run(os.Args[2])

	default:
		Usage()
	}
}
