package model

import "aliyuncs/ecs"

type DescribeDisksResponse struct {
	ecs.Response
	RegionId string
	TotalCount string
	PageNumber string
	PageSize string
	Disks DiskItemSetType
}

type DiskItemSetType struct {
	Disk []DiskItemType
}

type DiskItemType struct {
	DiskId string
	RegionId string
	ZoneId string
	DiskName string
	Description string
	Type string
	Category string
	Size int
	ImageId string
	SourceSnapshotId string
	ProductCode string
	Portable bool
	Status string
	OperationLocks OperationLocksType
	InstanceId string
	Device string
	DeleteWithInstance bool
	DeleteAutoSnapshot bool
	EnableAutoSnapshot bool
	CreationTime string
	AttachedTime string
	DetachedTime string
	DiskChargeType string
}