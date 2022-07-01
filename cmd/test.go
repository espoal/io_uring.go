package main

import (
	"bytes"
	"fmt"
	"io_uring.go/libs/liburing"
	"sync"
)

func main() {
	err := liburing.Init()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer liburing.Cleanup()

	go func() {
		fmt.Println("Func! ")
		for err := range liburing.Err() {
			fmt.Println(err)
		}
	}()

	var wg sync.WaitGroup
	// Read a file.
	err = liburing.ReadFile("../data/read.txt", func(buf []byte) {
		fmt.Println("Read! ")
		fmt.Println(bytes.NewBuffer(buf).String())
		//defer wg.Done()
		// handle buf
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Write something
	err = liburing.WriteFile("../data/write.txt", []byte("hello world"), 0644, func(n int) {
		fmt.Println("Write! ")
		//defer wg.Done()
		// handle n
	})
	// Call Poll to let the kernel know to read the entries.
	liburing.Poll()
	// Wait till all callbacks are done.
	wg.Wait()
}
