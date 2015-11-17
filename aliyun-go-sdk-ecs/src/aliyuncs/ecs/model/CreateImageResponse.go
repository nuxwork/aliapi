package model

import "aliyuncs/ecs"

type CreateImageResponse struct {
	ecs.Response
	ImageId string
}