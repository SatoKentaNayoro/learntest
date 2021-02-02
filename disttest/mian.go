package main

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"time"
)

func test() {
	var done = make(chan struct{}, 1)

	timeout, _ := context.WithTimeout(context.Background(), time.Second*10)
	//cancelCtx, cancelFunc := context.WithCancel(context.Background())
	go func() {
		time.Sleep(time.Second * 11)
		select {
		case <- timeout.Done():
			fmt.Println("goroutine 终止了")
		default:
			fmt.Println(done)
			done <- struct{}{}
		}

	}()
	select {
	case <-timeout.Done():
		fmt.Println("timeout")
	case <-done:
		fmt.Println("test done")
	}
}

func main() {
	buffer := bytes.NewBuffer([]byte{})
	timeOut := time.Second * 10
	done := make(chan struct{},1)
	go func() {
		post, err := http.Post("http://192.168.88.123:17801", "", buffer)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(post)
		done <- struct {}{}
	}()
	select {
	case <- time.After(timeOut):
		fmt.Println("time out")
	case <- done:
		fmt.Println("ok")
	}
}
