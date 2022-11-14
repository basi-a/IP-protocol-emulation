package main

import (
	iphdr "IPv4SoftwarePkg/ipheader"
	swpg "IPv4SoftwarePkg/softwarePkg"
	"fmt"
	"log"
)

/*发送端*/
func send()(*iphdr.IpHeader, []uint8)  {
	native_addr := "192.168.56.111"
	gateway_addr := "192.168.56.1"
	subnet_mask := "255.255.255.0"
	src_addr := native_addr
	dest_addr := "192.168.56.112"
	MTU := 1500 //MTU大小
	native_addr_uint32 := swpg.InetAtoN(native_addr)
	gateway_addr_uint32 := swpg.InetAtoN(gateway_addr)
	src_addr_uint32 := swpg.InetAtoN(src_addr)
	dest_addr_uint32 := swpg.InetAtoN(dest_addr)

	//发送端信息
	log.Println("发送端信息")
	log.Println("\tIPv4: ", native_addr)
	log.Println("\t默认网关: ", gateway_addr)
	log.Println("\t子网掩码: ", subnet_mask)
	//数据封装IP数据报
	header, data := swpg.IpAddingModule(src_addr_uint32, dest_addr_uint32)
	//处理模块处理数据报
	header, data = swpg.IpProcessingModule(header, data, native_addr_uint32, gateway_addr_uint32, MTU)
	return header, data
}
/*接收端*/
func receive(header *iphdr.IpHeader, data []uint8)  {
	native_addr := "192.168.56.112"
	gateway_addr := "192.168.56.1"
	subnet_mask := "255.255.255.0"
	MTU := 1500 //MTU大小
	native_addr_uint32 := swpg.InetAtoN(native_addr)
	gateway_addr_uint32 := swpg.InetAtoN(gateway_addr)

	//接收端信息
	log.Println("接收端信息")
	log.Println("\tIPv4: ", native_addr)
	log.Println("\t默认网关: ", gateway_addr)
	log.Println("\t子网掩码: ", subnet_mask)

	//处理模块处理接收到的数据报
	header, data = swpg.IpProcessingModule(header, data, native_addr_uint32, gateway_addr_uint32, MTU)

	headerbytes := swpg.StructToBytes(header)
	recheck := swpg.IPv4reCheckSum(headerbytes)
	//数据报中目的地址不是本机, 或数据报损坏(recheck=0xFFFF), 丢弃数据报
	if header.DestinationAddress == native_addr_uint32 && recheck == 0xFFFF {
		log.Println("数据报中目的地址是本机, 且数据报未损坏")
		log.Println("接收到的数据报:")
		log.Println("首部:",headerbytes)
		log.Println("数据:",data)
	}else{
		log.Println("数据报中目的地址不是本机, 或数据报损坏, 丢弃数据报")
	}
	
}

func main()  {
	log.Println("+++++++++++++++IP协议模拟程序  开始模拟+++++++++++++++")
	fmt.Println()
	log.Println("发送端发送数据报......")
	header, data := send()
	fmt.Println()
	log.Println("接收端接收数据报......")
	receive(header, data)
}