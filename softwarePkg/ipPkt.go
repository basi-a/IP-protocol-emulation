package softwarepkg

import (
	iphdr "IPv4SoftwarePkg/ipheader"
	"encoding/binary"
	"log"
)

/*
数据封装进IP数据报 data: yyz -> 0x79797a
*/
func IpPkt(src_addr uint32, dest_addr uint32) (*iphdr.IpHeader, []uint8){
	
	var IP_VER_HLEN 	uint8 = 0x45		/*定义IP首部的版本和首部长度, IPv4,首部长度20,对应0x45*/
	var IP_FRAGEMENT 	uint16 = 0x4000 	/*不分片*/
	var IP_TTL 			uint8 = 0x80		/*TTL=128*/
	var IP_PROTOCOL 	uint8 = 0x01		/*上层协议类型, 此处0x01表示ICMP*/
	/*为IP首部各字段赋值*/
	ipheader := iphdr.IpHeader{
		VerHlen: IP_VER_HLEN,
		Service: 0x00,
		Length: 32,
		Identification: 0xffff,
		FlagsAndFragementOffset: IP_FRAGEMENT,
		TimeToLive: IP_TTL,
		Protocol: IP_PROTOCOL,
		HeaderChecksum: 0x0000,
		SourceAddress: src_addr,
		DestinationAddress: dest_addr,
	}
	IP_HEADER_LEN := binary.Size(ipheader)	/*获取首部长度*/
	IP_DATA_LEN := ipheader.Length-uint16(IP_HEADER_LEN)//32-20=12
	var data = make([]uint8, IP_DATA_LEN)	
	//hex(yyz) : 0x79797a
	data[0] = 0x79
	data[1] = 0x79
	data[2] = 0x7a
	log.Println("初始校验和:",ipheader.HeaderChecksum)
	log.Println("数据:",data)
	return &ipheader, data
}
