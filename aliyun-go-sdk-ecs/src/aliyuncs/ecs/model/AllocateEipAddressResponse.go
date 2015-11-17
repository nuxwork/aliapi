package model

import "aliyuncs/ecs"

type AllocateEipAddressResponse struct {
	ecs.Response
	EipAddress string
	AllocationId string
}