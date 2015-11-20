package v20140526

import "aliyuncs/core"

type DescribeInstanceTypesRequest struct {
	core.Request
}

func NewDescribeInstanceTypesRequest() *DescribeInstanceTypesRequest {
	r := new(DescribeInstanceTypesRequest)
	r.Init("Ecs", "2014-05-26", "DescribeInstanceTypes")
	return r
}