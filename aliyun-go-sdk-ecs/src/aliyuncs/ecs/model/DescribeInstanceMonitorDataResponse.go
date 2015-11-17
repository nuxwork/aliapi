package model

import "aliyuncs/ecs"

type DescribeInstanceMonitorDataResponse struct {
	ecs.Response
	MonitorData MonitorDataType
}

type MonitorDataType struct {
	InstanceMonitorData []InstanceMonitorDataType
}

type InstanceMonitorDataType struct {
	InstanceId	string  //实例ID
	CPU	int  //CPU的使用比例，单位：百分比（%）
	IntranetRX	int  //云服务器实例接收到的数据流量，单位：kbits
	IntranetTX	int  //云服务器实例接发送的数据流量，单位：kbits
	IntranetBandwidth	int  //云服务器实例的带宽（单位时间内的网络流量），单位为kbits/s
	InternetRX	int  //云服务器实例接收到的数据流量，单位：kbits
	InternetTX	int  //云服务器实例接发送的数据流量，单位：kbits
	InternetBandwidth	int  //云服务器实例的带宽（单位时间内的网络流量），单位为kbits/s
	IOPSRead	int  //系统盘IO读操作，单位：次/s
	IOPSWrite	int  //系统盘IO写操作，单位：次/s
	BPSRead	int  //系统盘磁盘读带宽，单位：Byte/s
	BPSWrite	int  //系统盘磁盘写带宽，单位：Byte/s
	TimeStamp	string  //查询流量的时间点，按照ISO8601标准表示，并需要使用UTC时间。格式为：YYYY-MM-DDThh:mmZ
}