package softwarepkg

import (
	iphdr "IPv4SoftwarePkg/ipheader"
	"log"

)

/*
首部添加模块
*/
func IpAddingModule(src_addr uint32, dest_addr uint32) (*iphdr.IpHeader, []uint8) {
	/*数据封装进IP数据报*/
	log.Println("开始将数据封装到数据报......")
	header, data := IpPkt(src_addr, dest_addr)

	log.Println("IP首部及数据已经封装到数据报!!!")
	headerbytes := StructToBytes(header)
	log.Println("初始校验和:", header.HeaderChecksum)
	log.Println("首部:", headerbytes)
	log.Println("数据:", data)

	/*计算首部校验和, 然后赋值给首部校验和*/
	log.Println("计算首部校验和......")
	checkSum := IPv4CheckSum(headerbytes)
	header.HeaderChecksum = checkSum
	headerbytes = StructToBytes(header)
	log.Println("首部校验和:", header.HeaderChecksum)
	log.Println("首部:", headerbytes)
	return header, data
}

/*
处理模块, 当目的地址与本地地址相匹配时, 将数据报发送到重装模块, 否则发送到转发模块
*/
func IpProcessingModule(header *iphdr.IpHeader, data []uint8, native_addr uint32, gateway_addr uint32, MTU int) (*iphdr.IpHeader, []uint8)  {
	log.Println("处理模块开始处理......")
	log.Println("此网络MTU:", MTU)
	/*目的地址与本地地址相匹配时, 将数据报发送到重装模块*/
	if header.DestinationAddress == native_addr {
		log.Println("目的地址与本机地址相同, 数据报发送到重装模块...")
		data := IpReassemblyModule(header, data, MTU)
		return header, data
	}
	/*本机是路由器时TTL-1*/
	if native_addr == gateway_addr {
		log.Println("本机是路由器")
		header.TimeToLive = header.TimeToLive-1
	}
	/*TTL小于等于0时, 丢弃数据报, 发送ICMP差错报文*/
	if header.TimeToLive <= 0{
		log.Println("TTL<=0 丢弃数据报, 发送ICMP差错报文")
		return nil, nil
	}
	/*数据报发送到转发模块*/
	IpForwardingModule(header, data, MTU)
	log.Println("数据报已经发送到转发模块!!!")
	return nil, nil
}

/*
重装模块
*/
func IpReassemblyModule(header *iphdr.IpHeader, data []uint8, MTU int) []uint8 {
	/*分片偏移值是0且M也是0*/
	if header.FlagsAndFragementOffset == 0x4000 {
		return data
	}
	return nil
}

/*
转发模块
*/
func IpForwardingModule(header *iphdr.IpHeader, data []uint8, MTU int) {
	
}

/*
分片模块
*/
func IpFragmentationModule(header *iphdr.IpHeader, data []uint8, MTU int) {
	
}