package softwarepkg

import (
	iphdr "IPv4SoftwarePkg/ipheader"
	"bytes"
	"encoding/binary"
	"log"

)

/*
IPv4首部校验和计算
*/
func IPv4CheckSum(data []byte) uint16 {
	
	var sum uint32
	var length int = len(data)
	var index int
	/*以每16位为单位进行求和, 直到所有的字节全部求完或者只剩下一个8位字节(如果剩余一个8位字节说明字节数为奇数个)*/
	for length > 1 {
		sum += uint32(data[index])<<8 + uint32(data[index+1])
		index += 2
		length -= 2
	}
	/*如果字节数为奇数个, 要加上最后剩下的那个8位字节*/
	if length > 0 {
		sum += uint32(data[index])
	}
	/*加上高16位进位的部分*/
	sum += (sum >> 16)
	/*返回的时候求反*/
	return uint16(^sum)
}
/*
结构体转[]bype   type byte=uint8
*/
func StructToBytes(header *iphdr.IpHeader) []byte {
	buf := &bytes.Buffer{}
	err := binary.Write(buf, binary.BigEndian, header)
	if err != nil {
		panic(err)
	}
	headerbytes := buf.Bytes()
	return headerbytes
}


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
	checkSum := IPv4CheckSum(headerbytes)
	header.HeaderChecksum = checkSum
	headerbytes = StructToBytes(header)
	log.Println("计算得到的首部校验和:", header.HeaderChecksum)
	log.Println("计算完校验和之后的首部:", headerbytes)
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