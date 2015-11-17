package model

import "aliyuncs/ecs"

type DescribeInstanceVncUrlResponse struct {
	ecs.Response
	VncUrl string
}