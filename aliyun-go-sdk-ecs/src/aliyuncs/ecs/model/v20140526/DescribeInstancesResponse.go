package v20140526

import "aliyuncs/core"

type DescribeInstancesResponse struct {
	core.Response
	TotalCount int
	PageNumber int
	PageSize   int
	Instances  InstanceAttributesTypes
}

type InstanceAttributesTypes struct {
	Instance []InstanceAttributesType
}

type InstanceAttributesType struct {
	InstanceId              string
	InstanceName            string
	Description             string
	ImageId                 string
	RegionId                string
	ZoneId                  string
	InstanceType            string
	HostName                string
	SerialNumber            string
	ClusterId               string
	Status                  string
	VlanId                  string
	SecurityGroupIds        SecurityGroupIdSetType
	PublicIpAddress         IpAddressSetType
	InternetMaxBandwidthIn  int
	InternetMaxBandwidthOut int
	InternetChargeType      string
	CreationTime            string
	VpcAttributes           VpcAttributesType
	EipAddress              EipAddressAssociateType
	InnerIpAddress          IpAddressSetType
	InstanceNetworkType     string
	OperationLocks          OperationLocksType
	InstanceChargeType      string
	DeviceAvailable         bool
	IoOptimized             bool
	ExpiredTime             string
	Tags                    TagType
}

type SecurityGroupIdSetType struct {
	SecurityGroupId []string
}

type VpcAttributesType struct {
	VpcId            string
	VSwitchId        string
	PrivateIpAddress IpAddressSetType
	NatIpAddress     string
}

type IpAddressSetType struct {
	IpAddress []string
}

type EipAddressAssociateType struct {
	AllocationId       string
	IpAddress          string
	Bandwidth          int
	InternetChargeType string
}

type OperationLocksType struct {
	LockReason []string
}

type TagType struct {
	Tag []string
}
