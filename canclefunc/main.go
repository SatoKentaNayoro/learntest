package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var sig = make(chan os.Signal, 2)
	defer fmt.Println(1111)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-sig:
	}
}
