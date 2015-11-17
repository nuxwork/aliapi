package model

import "aliyuncs/ecs"

type AllocatePublicIpAddressResponse struct {
	ecs.Response
	IpAddress string
}