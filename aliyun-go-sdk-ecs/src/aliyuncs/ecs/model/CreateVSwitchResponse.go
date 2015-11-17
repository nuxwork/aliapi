package model

import "aliyuncs/ecs"

type CreateVSwitchResponse struct {
	ecs.Response
	VSwitchId string
}