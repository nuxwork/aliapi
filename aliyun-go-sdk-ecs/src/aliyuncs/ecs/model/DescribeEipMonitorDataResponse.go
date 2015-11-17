package model

import "aliyuncs/ecs"

type DescribeEipMonitorDataResponse struct {
  ecs.Response
  EipMonitorDatas  EipMonitorDataSetType  //实例的监控数据 EipMonitorDataType 数据集合。
}

type EipMonitorDataSetType struct {
  EipMonitorData []EipMonitorDataType
}

type EipMonitorDataType struct {
  EipRX  int  //一段时间（Period）内，EIP接收到的数据流量，单位：bytes。
  EipTX  int  //一段时间（Period）内，EIP接发送的数据流量，单位：bytes。
  EipFlow  int  //一段时间（Period）内，EIP网络流量，单位bytes。
  EipBandwidth  int  //弹性公网IP的带宽（单位时间内的网络流量），单位为bytes/s。
  EipPackets  int  //一段时间（Period）内，EIP接受和发送的报文总数。
  TimeStamp  string  //查询流量的时间点，按照ISO8601标准表示，并需要使用UTC时间。格式为：YYYY-MM-DDThh:mmZ
}