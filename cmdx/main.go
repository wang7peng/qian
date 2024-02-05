package main

// This command will be more big, more full, and more colorful

import (
	"os"
	Demo "qianDev/guidebook"
	L "qianDev/internal/app_big"
	F "qianDev/internal/app_small"
)

func init() {

}

func main() {
	if len(os.Args) == 1 {
		L.Usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "--version", "version", "-v", "-V":
		L.ShowVersion()

	case "--help", "help", "-h":
		L.Usage()

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
			L.ShowWarn()
		}

	case "--page", "page", "-p", "-P":
		// relate page demo
		Demo.Run(os.Args[2])

	default:
		L.Usage()
	}
}
