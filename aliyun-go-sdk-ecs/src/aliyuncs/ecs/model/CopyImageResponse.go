package model

import "aliyuncs/ecs"

type CopyImageResponse struct {
	ecs.Response
	ImageId string
}