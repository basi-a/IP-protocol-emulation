package softwarepkg
import (
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
