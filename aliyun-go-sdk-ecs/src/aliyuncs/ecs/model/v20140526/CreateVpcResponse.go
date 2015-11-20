package v20140526

import "aliyuncs/core"

type CreateVpcResponse struct {
	core.Response
	VpcId        string //系统分配的专有网络 ID
	VRouterId    string //路由器 Id
	RouteTableId string //路由表 Id
}
