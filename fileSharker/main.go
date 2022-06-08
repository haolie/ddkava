package main

import (
	"sync"

	_ "fileSharker/src"
	"fileSharker/src/Server"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	Server.Start(&wg)
	wg.Wait()
}
