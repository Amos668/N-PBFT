package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"time"
)

func Router(resp http.ResponseWriter, request *http.Request) {
	msg := request.FormValue("msg")
	r := new(Request)
	r.Timestamp = time.Now().UnixNano()
	r.ClientAddr = clientAddr
	r.Message = Message{
		msg,
		getRandom(),
	}
	br, err := json.Marshal(r)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(string(br))
	content := jointMessage(cRequest, br)
	tcpDial(content, nodeTable["P0"])
}

func clientSendMessageAndListen() {
	//开启客户端的本地监听（主要用来接收节点的reply信息）
	http.HandleFunc("/", Router)

	go http.ListenAndServe("127.0.0.1:8888", nil)
	fmt.Println(" ---------------------------------------------------------------------------------")
	fmt.Printf("客户端开启监听，地址：%s\n", clientAddr)
	fmt.Println(" ---------------------------------------------------------------------------------")
}

//返回一个十位数的随机数，作为msgid
func getRandom() int {
	x := big.NewInt(10000000000)
	for {
		result, err := rand.Int(rand.Reader, x)
		if err != nil {
			log.Panic(err)
		}
		if result.Int64() > 1000000000 {
			return int(result.Int64())
		}
	}
}
