package ipheader

/*
根据IP首部定义结构体
*/
type IpHeader struct {
	VerHlen					uint8		/*IP首部的版本和首部长度, 各占4位*/
	Service					uint8		/*服务类型*/
	Length 					uint16		/*数据报总长度（数据长度+hlen字段*4）*/
	Identification			uint16		/*标识*/
	FlagsAndFragementOffset	uint16		/*前3位为标志，后13位为分片偏移*/
	TimeToLive				uint8		/*生存时间*/
	Protocol				uint8		/*协议, 定义了上层协议*/
	HeaderChecksum			uint16		/*首部校验和*/
	SourceAddress			uint32		/*源IP地址*/	
	DestinationAddress		uint32		/*目的IP地址*/	
}
