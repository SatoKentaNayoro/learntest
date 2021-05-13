/**
 _*_ @Author: IronHuang _*_
 _*_ @blog:https://www.dvpos.com/ _*_
 _*_ @Date: 2021/5/8 上午10:50 _*_
**/

package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"time"
)

func main() {
	db, err := bolt.Open("/Users/huangxuchen/boltDBtest/test.db", 0600, &bolt.Options{Timeout: time.Second * 20})
	if err != nil {
		log.Fatal(err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("tasks"))
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte("tasks"))
			if err != nil {
				log.Fatal(err)
			}
		}
		bucket.Put([]byte("222"), []byte("asdafasfasfasfa"))
		bucket.Put([]byte("333"), []byte("ssss"))
		bucket.Put([]byte("asd"), []byte("asddd"))

		fmt.Println(bucket.Stats().KeyN)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
