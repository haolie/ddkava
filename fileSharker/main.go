package main

import (
	"fmt"

	"fileSharker/config"
	"fileSharker/src/fileScan"
)

func main() {
	config.Load()

	for _, pth := range config.GetPaths() {
		list, err := fileScan.ScanDir(pth, true)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, item := range list {
			fmt.Printf("fileName:%s key:%s", item.FileName, item.FileKey)
		}
	}
}
