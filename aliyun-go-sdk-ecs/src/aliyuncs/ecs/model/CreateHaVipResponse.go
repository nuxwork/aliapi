package model

import "aliyuncs/ecs"

type CreateHaVipResponse struct {
	ecs.Response
	HaVipId string
}