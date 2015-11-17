package model

import "aliyuncs/ecs"

type DescribeInstanceTypesRequest struct {
	ecs.Request
}

func NewDescribeInstanceTypesRequest() *DescribeInstanceTypesRequest{
	r := new(DescribeInstanceTypesRequest)
	r.Set("Action", "DescribeInstanceTypes")
	return r
}
