package model

import "aliyuncs/ecs"

type DescribeRegionsRequest struct {
	ecs.Request
}

func NewDescribeRegionsRequest() *DescribeRegionsRequest{
	r := new(DescribeRegionsRequest)
	r.Set("Action", "DescribeRegions")
	return r
}
