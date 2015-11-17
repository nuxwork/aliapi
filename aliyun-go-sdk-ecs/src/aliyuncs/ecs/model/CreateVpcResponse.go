package model

import "aliyuncs/ecs"

type CreateVpcResponse struct {
	ecs.Response
	VpcId	string  //系统分配的专有网络 ID
	VRouterId	string  //路由器 Id
	RouteTableId	string  //路由表 Id
}