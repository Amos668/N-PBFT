package main

import (
	"log"
	"os"
)

const nodeCount = 4

//客户端的监听地址
var clientAddr = "127.0.0.1:8888"

//节点池，主要用来存储监听地址
var nodeTable map[string]string

func main() {
	//为四个节点生成公私钥
	genRsaKeys()
	nodeTable = map[string]string{
		"P0": "127.0.0.1:8000",
		"P1": "127.0.0.1:8001",
		"P2": "127.0.0.1:8002",
		"P3": "127.0.0.1:8003",
		"P4": "127.0.0.1:8084",
		//"P5": "127.0.0.1:8005",
		//"P6": "127.0.0.1:8006",
		//"P7":  "127.0.0.1:8007",
		//"P8":  "127.0.0.1:8008",
		//"P9":  "127.0.0.1:8009",
		//"P10": "127.0.0.1:8010",
		//"P11": "127.0.0.1:8011",
		//"P12": "127.0.0.1:8012",
		//"P13": "127.0.0.1:8013",
		//"P14": "127.0.0.1:8014",
		//"P15": "127.0.0.1:8015",
	}

	if len(os.Args) != 2 {
		log.Panic("输入的参数有误！")
	}

	nodeID := os.Args[1]
	if nodeID == "client" {
		clientSendMessageAndListen() //启动客户端程序
	} else if addr, ok := nodeTable[nodeID]; ok {
		p := NewPBFT(nodeID, addr)
		go p.tcpListen() //启动节点
	} else {
		log.Fatal("无此节点编号！")
	}
	select {}
}
