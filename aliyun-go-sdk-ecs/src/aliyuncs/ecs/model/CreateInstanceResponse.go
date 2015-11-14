package model

import "aliyuncs/ecs"

type CreateInstanceResponse struct {
	ecs.Response
	InstanceId string
}