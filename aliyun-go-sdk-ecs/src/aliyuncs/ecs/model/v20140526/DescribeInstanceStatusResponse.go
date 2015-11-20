package v20140526

import "aliyuncs/core"

type DescribeInstanceStatusResponse struct {
	core.Response
	TotalCount       int
	PageNumber       int
	PageSize         int
	InstanceStatuses InstanceStatusSetType
}

type InstanceStatusSetType struct {
	InstanceStatus []InstanceStatusItemType
}

type InstanceStatusItemType struct {
	InstanceId string
	Status     string
}
