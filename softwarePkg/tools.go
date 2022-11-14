package softwarepkg

import (
	iphdr "IPv4SoftwarePkg/ipheader"
	"bytes"
	"encoding/binary"
	"fmt"
	"math/big"
	"net"
)

/*
uint32 转字符串IP地址
*/
func InetNtoA(ip uint32) string {
    return fmt.Sprintf("%d.%d.%d.%d",byte(ip>>24), byte(ip>>16), byte(ip>>8), byte(ip))
}

/*
字符串IP地址转 uint32
*/
func InetAtoN(ip string) uint32 {
    ret := big.NewInt(0)
    ret.SetBytes(net.ParseIP(ip).To4())
    return uint32(ret.Uint64())
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
/*
校验和接收端校验计算
*/
func IPv4reCheckSum(data []byte) uint16 {
	
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
	return uint16(sum)
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