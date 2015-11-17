package model

import "aliyuncs/ecs"

type CreateDiskResponse struct {
	ecs.Response
	DiskId string
}