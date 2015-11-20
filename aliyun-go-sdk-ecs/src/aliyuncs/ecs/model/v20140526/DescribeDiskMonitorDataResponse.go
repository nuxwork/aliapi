package v20140526

import "aliyuncs/core"

type DescribeDiskMonitorDataResponse struct {
	core.Response
	TotalCount  int                    //磁盘监控数据的总个数
	MonitorData DiskMonitorDataSetType //磁盘的监控数据 DiskMonitorDataType 组成的集合。
}

type DiskMonitorDataSetType struct {
	DiskMonitorData []DiskMonitorDataType
}

type DiskMonitorDataType struct {
	DiskId    string //磁盘编号
	IOPSRead  int    //磁盘IO读操作，单位：次/s
	IOPSWrite int    //磁盘IO写操作，单位：次/s
	IOPSTotal int    //磁盘IO读写总操作，单位：次/s
	BPSRead   int    //磁盘读带宽，单位：byte/s
	BPSWrite  int    //磁盘写带宽，单位：byte/s
	BPSTotal  int    //磁盘读写总带宽，单位：byte/s
	TimeStamp string //查询的时间点，按照ISO8601标准表示，并需要使用UTC时间。格式为：YYYY-MM-DDThh:mmZ
}
