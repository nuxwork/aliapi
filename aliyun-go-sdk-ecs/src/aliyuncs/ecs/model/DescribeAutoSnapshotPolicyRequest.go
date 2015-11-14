package model

import "aliyuncs/ecs"

type DescribeAutoSnapshotPolicyRequest struct {
	ecs.Request
}

func NewDescribeAutoSnapshotPolicyRequest() *DescribeAutoSnapshotPolicyRequest{
	r := new(DescribeAutoSnapshotPolicyRequest)
	r.Set("Action", "DescribeAutoSnapshotPolicy")
	return r
}