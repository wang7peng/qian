package main

import (
	"os"
	F "qianDev/internal/app_small"
)

func init() {

}

func main() {
	if len(os.Args) == 1 {
		F.Usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "--version", "-v", "-V":
		F.ShowVersion()

	case "--help", "-h":
		F.Usage()

	case "--create", "-c":
		// 创建分机
		if len(os.Args) > 2 {
			F.CreateExtension(os.Args[2])
		} else {
			F.ShowWarn()
		}

	case "--delete", "-d":
		// 删除分机
		if len(os.Args) > 2 {
			F.DeleteExtension(os.Args[2])
		} else {
			F.ShowWarn()
		}

	default:
		F.Usage()
	}
}
