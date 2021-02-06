/**
 _*_ @Author: IronHuang _*_
 _*_ @blog:https://www.dvpos.com/ _*_
 _*_ @Date: 2021/2/6 下午12:50 _*_
**/

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	fileContent, err := os.Open("computeTimeCost/7714p1costtime")
	if err != nil {
		fmt.Println("error:", err.Error())
		return
	}
	defer fileContent.Close()
	br := bufio.NewReader(fileContent)
	contentSlice := make([]string, 0)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		if sl := string(line); sl != "--" {
			contentSlice = append(contentSlice, sl)
		}
	}
	for i, j, k, layerNo := 0, 1, 2, 1; k < len(contentSlice); layerNo++ {
		TimeI, err := time.Parse("2006-01-02T15:04:05", contentSlice[i])
		if err != nil {
			log.Fatal(err)
		}
		TimeJ, err := time.Parse("2006-01-02T15:04:05", contentSlice[j])
		if err != nil {
			log.Fatal(err)
		}
		TimeK, err := time.Parse("2006-01-02T15:04:05", contentSlice[k])
		if err != nil {
			log.Fatal(err)
		}
		cmpCost := TimeJ.Sub(TimeI)
		storeCost := TimeK.Sub(TimeJ)
		fmt.Printf("layer %d 计算时间：%v，存储时间：%v\n", layerNo, cmpCost, storeCost)
		i += 3
		j += 3
		k += 3
	}

}
