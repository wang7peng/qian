package main

import (
	"fmt"
	"os"
	Demo "qianDev/guidebook"
)

func init() {

}

func Usage() {
	fmt.Println("Usage")
	fmt.Println("  -v      查看当前版本")
	fmt.Println("  -c      新建PJSIP话机终端")
	fmt.Println("  -p [n]  提供书页, 运行对应页的案例")
	fmt.Println("  -h      使用帮助")
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

	case "-c":
		CreateEndPoint()

	case "--page", "page", "-p", "-P":
		// relate page demo
		Demo.Run(os.Args[2])

	default:

	}
}
