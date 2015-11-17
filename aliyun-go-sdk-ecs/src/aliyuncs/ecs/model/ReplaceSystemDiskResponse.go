package model

import "aliyuncs/ecs"

type ReplaceSystemDiskResponse struct {
	ecs.Response
	DiskId string
}