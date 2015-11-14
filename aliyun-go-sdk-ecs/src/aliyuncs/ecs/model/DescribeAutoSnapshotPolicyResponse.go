package model

import "aliyuncs/ecs"

type DescribeAutoSnapshotPolicyResponse struct {
	ecs.Response
	AutoSnapshotExcutionStatus AutoSnapshotExecutionStatusType
	AutoSnapshotOccupation int
	AutoSnapshotPolicy AutoSnapshotPolicyType
}

type AutoSnapshotExecutionStatusType struct {
	DataDiskExcutionStatus string
	SystemDiskExcutionStatus string
}

type AutoSnapshotPolicyType struct {
	DataDiskPolicyRetentionLastWeek bool
	DataDiskPolicyRetentionDays int
	DataDiskPolicyEnabled bool
	SystemDiskPolicyEnabled bool
	SystemDiskPolicyRetentionLastWeek bool
	DataDiskPolicyTimePeriod int
	SystemDiskPolicyTimePeriod int
	SystemDiskPolicyRetentionDays int
}


