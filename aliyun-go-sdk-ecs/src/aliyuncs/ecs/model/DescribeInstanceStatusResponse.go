package model

import "aliyuncs/ecs"

type DescribeInstanceStatusResponse struct {
	ecs.Response
	TotalCount int
	PageNumber int
	PageSize int
	InstanceStatuses InstanceStatusSetType
}

type InstanceStatusSetType struct {
	InstanceStatus []InstanceStatusItemType
}

type InstanceStatusItemType struct {
	InstanceId string
	Status string
}