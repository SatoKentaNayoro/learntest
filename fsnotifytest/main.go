/**
 _*_ @Author: IronHuang _*_
 _*_ @github:https://github.com/hxuchen _*_
 _*_ @Date: 2021/6/25 上午10:55 _*_
**/

package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	err = watcher.Add("/Users/huangxuchen/gotest_pro/learntest/fsnotifytest")
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case a := <-watcher.Events:
			fmt.Println(a.Name)
			fmt.Println(a.String())
			fmt.Println(a.Op)
		}
	}
}
