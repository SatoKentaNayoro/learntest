/**
 _*_ @Author: IronHuang _*_
 _*_ @github:https://github.com/hxuchen _*_
 _*_ @Date: 2021/6/27 下午2:41 _*_
**/

package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type Test struct {
	A string `bson:"a"`
	B string `bson:"b"`
	C string `bson:"c"`
}

type UpdateModel struct {
	A string
	B string
	C string

	FA string
	FB string
	FC string
}

func main() {
	//t := Test{
	//	"aaa111",
	//	"bbb111",
	//	"ccc111",
	//}

	ut := UpdateModel{
		A:  "aaa",
		C:  "sdadada",
		FA: "aaa",
	}

	ctx := context.Background()
	// initialize mongodb
	//uri := fmt.Sprintf("mongodb://127.0.0.1:27017/?authSource=admin&authMechanism=SCRAM-SHA-1&ssl=true")
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb")

	// 连接到MongoDB
	mgoCli, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalln("failed to connect with error:", err)
	}
	collection := mgoCli.Database("live").Collection("live")
	//b:= bson.D{
	//	{
	//		"a",
	//		t.A,
	//	},
	//	{
	//		"b",
	//		t.B,
	//	},
	//	{
	//		"c",
	//		t.C,
	//	},
	//}
	//_, err = collection.InsertOne(ctx, b)
	//if err != nil {
	//	log.Fatal(err)
	//}

	f1Base := bson.D{
		{
			"a",
			ut.FA,
		},
		{
			"b",
			ut.FB,
		},
		{
			"c",
			ut.FC,
		},
	}

	var f1 bson.D
	// del empty param
	for idx, v := range f1Base {
		if v.Key == "start" || v.Key == "end" {
			if v.Value == nil {
				if idx < len(f1Base)-1 {
					f1 = append(f1Base[:idx], f1Base[idx+1:]...)
				} else {
					f1 = f1Base[:idx-1]
				}
			}
		} else {
			if v.Value == "" {
				if idx < len(f1Base)-1 {
					f1 = append(f1Base[:idx], f1Base[idx+1:]...)
				} else {
					f1 = f1Base[:idx-1]
				}
			}
		}
	}

	fmt.Println("f1:", f1)

	//u1 := bson.M{
	//	"$set":bson.M{
	//		"a":ut.A,
	//		"b":"",
	//	},
	//}
	uBase := bson.D{
		{
			"a",
			ut.A,
		},
		{
			"b",
			ut.B,
		},
		{
			"c",
			ut.C,
		},
	}
	var u bson.D
	// del empty param
	for idx, v := range uBase {
		if v.Key == "start" || v.Key == "end" {
			if v.Value == nil {
				if idx < len(uBase)-1 {
					u = append(uBase[:idx], uBase[idx+1:]...)
				} else {
					u = uBase[:idx-1]
				}
			}
		} else {
			if v.Value == "" {
				if idx < len(uBase)-1 {
					u = append(uBase[:idx], uBase[idx+1:]...)
				} else {
					u = uBase[:idx-1]
				}
			}
		}
	}

	u1 := bson.D{
		{
			"$set",
			u,
		},
	}
	fmt.Println(u1)

	many, err := collection.UpdateMany(ctx, f1, u1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(many)
	timeN := time.Time{}
	fmt.Println(timeN.IsZero())
}
