package v20140526

import "aliyuncs/core"

type DescribeAutoSnapshotPolicyRequest struct {
	core.Request
}

func NewDescribeAutoSnapshotPolicyRequest() *DescribeAutoSnapshotPolicyRequest {
	r := new(DescribeAutoSnapshotPolicyRequest)
	r.Init("Ecs", "2014-05-26", "DescribeAutoSnapshotPolicy")
	return r
}
