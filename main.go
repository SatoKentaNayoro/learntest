package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type TestRpc struct{}

func (t *TestRpc) Add(i int) int {
	i++
	return i
}

func main() {
	//var fileList = make([]string,0)
	//files, err := GetAllFile("/Users/huangxuchen/gotest_pro/learntest",fileList)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(files)
	//stat, _ := os.Stat("/22")
	//fmt.Println(stat)
	//buf := make([]byte,1)
	//open, _ := os.Open("11.txt")
	//seek, _ := open.Seek(4, 0)
	//open.Read(buf)
	//fmt.Println(seek)

	//var t = new(crc32.Table)
	//file,_ := os.Open("22.txt")
	//reader := bufio.NewReader(file)
	//all, _ := ioutil.ReadAll(reader)
	////checksum := crc32.Checksum(all, t)
	////fmt.Println(checksum)
	////int32Buf := new(bytes.Buffer)
	////binary.Write(int32Buf, binary.LittleEndian, checksum)
	////fmt.Println(int32Buf.Bytes())
	//
	//ieee := crc32.NewIEEE()
	//io.WriteString(ieee,string(all))
	//s := ieee.Sum32()
	//fmt.Println(s)
	//size, err := GetAllFileSize("/Users/huangxuchen/学习资料" )
	//size, err := GetSrcSize("/Users/huangxuchen/学习资料")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(size)
	//a := "/22/dd/rr/aasdd/a/"
	//split := strings.Split(a, "/")
	//b := strings.Join(append(split[:len(split)-2], "cache"),"/")
	//fmt.Println(strings.TrimRight(a,"/"))
	//fmt.Println(b)
	//tMap := make(map[string]struct{})
	//tMap["2"] = struct{}{}
	//_, ok := tMap["2"]
	//fmt.Println(tMap)
	//fmt.Println(ok)
	//filepath.Walk("/Users/huangxuchen/gotest_pro/learntest/websocketClient", func(path string, info os.FileInfo, err error) error {
	//	if !info.Mode().IsRegular() {
	//		return nil
	//	}else {
	//		fmt.Println(path)
	//		fmt.Println(info.Name())
	//	}
	//	return nil
	//})
}

type ttt struct {
	a string
	b int
}

func (t *ttt) lll() {
	t.a = "2222"
}

func GetAllFile(src string, fileList []string) ([]string, error) {
	slash := filepath.FromSlash(src)
	dir, err := ioutil.ReadDir(slash)
	if err != nil {
		return fileList, err
	}
	for _, fi := range dir {
		// 如果还是文件夹
		if fi.IsDir() {
			fullDir := filepath.Join(slash, fi.Name())
			// 继续遍历
			fileList, err = GetAllFile(fullDir, fileList)
			if err != nil {
				return fileList, err
			}
		} else {
			fullName := filepath.Join(slash, fi.Name())
			fileList = append(fileList, fullName)
		}
	}
	return fileList, nil
}

func GetAllFileSize(src string) (uint64, error) {
	var sizeTotal uint64
	stat, err := os.Stat(src)
	if err != nil {
		return sizeTotal, err
	}
	if stat.IsDir() {
		slash := filepath.FromSlash(src)
		dir, err := ioutil.ReadDir(slash)
		if err != nil {
			return sizeTotal, err
		}
		for _, fi := range dir {
			// 如果还是文件夹
			if fi.IsDir() {
				fullDir := filepath.Join(slash, fi.Name())
				// 继续遍历
				_, err = GetAllFileSize(fullDir)
				if err != nil {
					return sizeTotal, err
				}
			} else {
				fmt.Println(fi.Name())
				fmt.Println(fi.Size())
				sizeTotal += uint64(fi.Size())
			}
		}
	}
	return sizeTotal, nil
}

func GetSrcSize(srcPath string) (uint64, error) {
	stat, err := os.Stat(srcPath)
	if err != nil {
		return 0, err
	}
	size := uint64(0)
	if stat.IsDir() {
		filepath.Walk(srcPath, func(path string, info os.FileInfo, err error) error {
			size += uint64(info.Size())
			return err
		})
	} else {
		size = uint64(stat.Size())
	}
	return size, nil
}
