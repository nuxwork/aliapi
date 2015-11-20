package v20140526

import "aliyuncs/core"

type DescribeRegionsRequest struct {
	core.Request
}

func NewDescribeRegionsRequest() *DescribeRegionsRequest {
	r := new(DescribeRegionsRequest)
	r.Init("Ecs", "2014-05-26", "DescribeRegions")
	return r
}
