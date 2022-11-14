package main

import (
	pc "IPv4SoftwarePkg/pcs"
	"fmt"
	"log"
)

func main()  {
	log.Println("+++++++++++++++IP协议模拟程序-开始模拟+++++++++++++++")
	fmt.Println()
	log.Println("发送端发送数据报......")
	header, data := pc.Send()
	fmt.Println()
	log.Println("接收端接收数据报......")
	pc.Receive(header, data)
}