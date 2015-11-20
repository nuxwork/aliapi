package v20140526

import "aliyuncs/core"

type DescribeVRoutersResponse struct {
	core.Response
	TotalCount int            //列表条条目数
	PageNumber int            //当前页码
	PageSize   int            //当前分页包含多少条目
	VRouters   VRouterSetType //VRouter 详情 VRouterSetType 组成的集合
}

type VRouterSetType struct {
	VRouter []VRouterType
}

type VRouterType struct {
	VRouterId     string //路由器的Id
	RegionId      string //地域Id
	VpcId         string //专有网络Id
	RouteTableIds string //虚拟路由表Id列表
	VRouterName   string //路由器名称
	Description   string //路由器的描述，不填则为空，默认值为空，[2,256]英文或中文字符，不能以http:// 和https:// 开头。
	CreationTime  string //创建时间。按照ISO8601标准表示，并需要使用UTC时间。格式为：YYYY-MM-DDThh:mmZ
}
