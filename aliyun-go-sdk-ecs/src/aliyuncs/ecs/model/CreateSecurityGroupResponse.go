package model

import "aliyuncs/ecs"

type CreateSecurityGroupResponse struct {
	ecs.Response
	SecurityGroupId string
}