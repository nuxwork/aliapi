package model

import "aliyuncs/ecs"

type DescribeImageSharePermissionResponse struct {
	ecs.Response
	ImageId string
	RegionId string
	TotalCount int
	PageNumber int
	PageSize int
	ShareGroups ShareGroupTypes
	Accounts AccountTypes
}

type ShareGroupTypes struct {
	ShareGroup []ShareGroupType
}

type ShareGroupType struct {
	Group string
}

type AccountTypes struct {
	Account []AccountType
}

type AccountType struct {
	AliyunId string
}