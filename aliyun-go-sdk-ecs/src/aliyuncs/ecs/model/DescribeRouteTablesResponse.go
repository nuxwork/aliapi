package model

import "aliyuncs/ecs"

type DescribeRouteTablesResponse struct {
	ecs.Response
	RouteTables	RouteTableSetType  //RouteTable 详情 RouteTableSetType 的集合
	TotalCount	int  //列表条条目数
	PageNumber	int  //当前页码
	PageSize	int  //当前分页包含多少条目
}

type RouteTableSetType struct {
	RouteTable []RouteTableType
}

type RouteTableType struct {
	VRouterId	string  //路由器的ID
	RouteTableId	string  //路由表的ID
	RouteEntrys	RouteEntrySetType  //路由条目详情RouteEntrySetType组成的集合
	RouteTableType	string  //路由表类型，System | Custom 二者选其一
	CreationTime	string  //创建时间。按照ISO8601标准表示，并需要使用UTC时间。格式为：YYYY-MM-DDThh:mmZ
}

type RouteEntrySetType struct {
	RouteEntry []RouteEntryType
}

type RouteEntryType struct {
	RouteTableId	string  //路由条目所在的路由器
	DestinationCidrBlock	string  //目标网段地址
	Type	string  //路由类型(System | Custom)
	NextHopType	string  //下一跳的目标对象类型
	NextHopId	string  //下一跳实例ID
	Status	string  //路由条目状态Pending | Available | Modifying
}
