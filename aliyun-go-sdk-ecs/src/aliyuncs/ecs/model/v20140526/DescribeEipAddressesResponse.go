package v20140526

import "aliyuncs/core"

type DescribeEipAddressesResponse struct {
	core.Response
	TotalCount   int
	PageNumber   int
	PageSize     int
	EipAddresses EipAddressSetType
}

type EipAddressSetType struct {
	RegionId           string
	IpAddress          string
	AllocationId       string
	Status             string
	InstanceType       string
	InstanceId         string
	Bandwidth          int
	InternetChargeType string
	OperationLocks     OperationLocksType
	AllocationTime     string
}
