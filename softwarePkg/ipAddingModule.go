package softwarepkg

import (
	iphdr "IPv4SoftwarePkg/ipheader"
	"bytes"
	"encoding/binary"
	"log"
)

/*
首部添加模块
*/
func IpAddingModule(src_addr uint32, dest_addr uint32) (*iphdr.IpHeader, []uint8) {
	/*数据封装进IP数据报*/
	log.Println("===开始将数据封装到数据报===")
	header, data := IpPkt(src_addr, dest_addr)
	log.Println("IP首部及数据已经封装到数据报!!!")
	/*结构体转[]bype   type byte=uint8 */
	buf := &bytes.Buffer{}
	err := binary.Write(buf, binary.BigEndian, header)
	if err != nil {
		panic(err)
	}
	/*计算首部校验和, 然后赋值给首部校验和*/
	checkSum := IPv4CheckSum(buf.Bytes())
	header.HeaderChecksum = checkSum
	log.Println("计算完得到的首部校验和:",header.HeaderChecksum)
	return header, data
}
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
