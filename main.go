package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

type TestRpc struct{}

func (t *TestRpc) Add(i int) int {
	i++
	return i
}

type aaa struct {
	c map[string]bbb
}

type bbb struct {
	ba []int
}

type YoungInfo struct {
	Name string
	Age  int
}

type SyncState struct {
	young     sync.Map
	youngSize int
	mu        sync.Mutex
}

func main() {
	//syncMap := sync.Map{}
	//for i := 0; i < 10; i++ {
	//	y := &YoungInfo{
	//		Name: string(i)+"a",
	//		Age: i,
	//	}
	//	store, loaded := syncMap.LoadOrStore(i, y)
	//}
	//syncMap.Range(func(key, value interface{}) bool {
	//	info := value.(*YoungInfo)
	//	go myPrint(info)
	//	time.Sleep(time.Second)
	//	return true
	//})
	//
	//time.Sleep(time.Second * 20)
	//myContext := context.WithValue(context.Background(),"a","b")
	//fmt.Println(myContext.Value("a"))
	getwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(getwd)
}

func Print(y *YoungInfo) {
	fmt.Println(*y)
}

var chan1 = make(chan struct{})

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
