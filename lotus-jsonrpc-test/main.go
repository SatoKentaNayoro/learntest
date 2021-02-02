package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type reqBody struct {
	Jsonrpc string        `json:"jsonrpc"`
	Id      int           `json:"id"`
	Params  []interface{} `json:"params"`
	Method  string        `json:"method"`
}

func NewReqBody(info []interface{}, method string) *reqBody {
	req := new(reqBody)
	req.Id = 1
	req.Params = info
	req.Jsonrpc = "2.0"
	req.Method = method
	return req
}

func (reqB *reqBody) getResByJsonRpc(address string) (back []byte, err error) {
	res, err1 := http.Post("http://"+address, "application/json", reqB.solvePostBody())
	if err1 != nil || res.StatusCode != http.StatusOK {
		if err1 == nil {
			err = errors.New(fmt.Sprintf("connection failed to %s,%s", address, res.Status))
		} else {
			err = err1
		}
		return
	}
	defer res.Body.Close()
	back, err = ioutil.ReadAll(res.Body)
	return
}

func (reqB *reqBody) solvePostBody() *bytes.Buffer {
	marshal, err1 := json.Marshal(*reqB)
	if err1 != nil {
		return nil
	}
	return bytes.NewBuffer(marshal)
}

type resBody struct {
	Jsonrpc string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
	Error   *resError    `json:"error"`
	Id      int         `json:"id"`
}

type resError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewResBody() *resBody {
	return new(resBody)
}

func (resB *resBody) UnMarshal(res []byte, model interface{}) (interface{}, error) {
	d1 := json.NewDecoder(bytes.NewReader(res))
	d1.UseNumber()
	err1 := d1.Decode(resB)
	if err1 != nil {
		return nil, err1
	}
	if resB.Error == nil {
		return nil,errors.New(resB.Error.Message)
	}
	marshal, err1 := json.Marshal(resB.Result)
	if err1 != nil {
		return nil, err1
	}
	d2 := json.NewDecoder(bytes.NewReader(marshal))
	err1 = d2.Decode(model)
	if err1 != nil {
		return nil, err1
	}
	return model, err1
}

type SectorID struct {
	Miner  uint64
	Number uint64
}

func main() {
	back, err := NewReqBody([]interface{}{1}, "WorkerJsonRpc.GetCurrentWorkingTask").getResByJsonRpc("192.168.88.172:18101")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var model = new(SectorID)
	_, err = NewResBody().UnMarshal(back, model)
	if err != nil {
		fmt.Println(err.Error())
		return
	}


}