package v20140526

import "aliyuncs/core"

type DescribeInstanceTypesResponse struct {
	core.Response
	InstanceTypes InstanceTypeItemSetType //由实例规格项 InstanceTypeItemType 组成的集合
}

type InstanceTypeItemSetType struct {
	InstanceType []InstanceTypeItemType
}

type InstanceTypeItemType struct {
	InstanceTypeId string //实例规格的ID
	CpuCoreCount   int    //CPU的内核数目
	MemorySize     uint64 //内存大小，单位GB
}
