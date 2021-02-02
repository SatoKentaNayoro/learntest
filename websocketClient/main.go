package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"time"
)

type Info struct {
	FileDir  string
	FileName string
	Data     []byte
	Done     bool
}

var limit = make(chan struct{}, 4)

func upload(num int) {
	u := url.URL{Scheme: "ws", Host: "127.0.0.1:17801", Path: "/upload"}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	defer conn.Close()
	done := make(chan struct{})

	const BUFFER_SIZE = 64 * 1024 * 1024
	buf := make([]byte, BUFFER_SIZE)

	source, err := os.Open("/Users/iron_huang/Downloads/阳光电影www.ygdy8.com.致允熙.BD.1080p.韩语中字.mkv/阳光电影www.ygdy8.com.致允熙.BD.1080p.韩语中字.mkv")
	if err != nil {
		fmt.Println("open file err:", err.Error())
		return
	}

	defer func() {
		err2 := source.Close()
		if err2 != nil && err == nil {
			err = err2
		}
	}()

Loop:
	for {
		var info = Info{
			FileDir:  "/Users/iron_huang",
			FileName: fmt.Sprintf("%d.mkv", num),
		}
		select {
		case <-interrupt:
			log.Println("interrupt")
			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			close(limit)
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		default:
			n, err := source.Read(buf)
			if err != nil && err != io.EOF {
				return
			}
			if n == 0 {
				info.Done = true
				break Loop
			}
			info.Data = buf[:n]
			marshal, err := json.Marshal(info)
			if err != nil {
				fmt.Println("json err:", err.Error())
			}
			err = conn.WriteMessage(websocket.TextMessage, marshal)
			if err != nil {
				fmt.Println("send file error:", err.Error())
				return
			}
			_, message, err := conn.ReadMessage()
			if mes := string(message); mes != "ok" || err != nil {
				log.Println("read:", err)
				return
			}
		}
	}
	fmt.Println("send over success")
	conn.Close()
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	for i := 0; i < 4; i++ {
		limit <- struct{}{}
	}
	fmt.Println(len(limit))
	i := 0
	for i < 4 {

		select {
		case <-limit:
			i++
			wg.Add(1)
			go upload(i)
		default:
		}
	}
	wg.Wait()
}
