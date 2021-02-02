package main

import (
	"fmt"
	"io"
	"os"
	"path"
)

func main() {
	//src := "/Users/iron_huang/IdeaProjects/tests/clienttest"
	//dst := "/Users/iron_huang/IdeaProjects/tests/2222"
	//filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
	//	targetPath := dst + "/"+info.Name()
	//	copy(path,targetPath)
	//
	//	return nil
	//})
	//dir := path.Dir("/Users/iron_huang/IdeaProjects/tests/2222")
	//fmt.Println(dir)
	////a := strings.Split("/Users/iron_huang/IdeaProjects/tests/2222",dir)
	//a :=strings.TrimPrefix("/Users/iron_huang/IdeaProjects/tests/2222",dir)
	a := path.Base("/Users/iron_huang/IdeaProjects/tests/2222")
	fmt.Println(a)
}


func copy(src, dst string) (err error) {

	if src == dst {
		return nil
	}
	const BUFFER_SIZE = 64 * 1024 * 1024
	buf := make([]byte, BUFFER_SIZE)

	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer func() {
		err2 := source.Close()
		if err2 != nil && err == nil {
			err = err2
		}
	}()

	MakeDirIfNotExists(path.Dir(dst))
	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func() {
		err2 := destination.Close()
		if err2 != nil && err == nil {
			err = err2
		}
	}()

	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		if _, err := destination.Write(buf[:n]); err != nil {
			return err
		}
	}

	return nil
}


func MakeDirIfNotExists(p string) error {

	// Check if parent dir exists. If not exists, create it.
	parentPath := path.Dir(p)

	_, err := os.Stat(parentPath)
	if err != nil && os.IsNotExist(err) {
		err = MakeDirIfNotExists(parentPath)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	// If parent dir exists. make dir.
	err = os.Mkdir(p, 0755)
	if err != nil && os.IsExist(err) {
		return nil
	} else {
		return err
	}
}
