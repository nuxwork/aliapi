package model

import "aliyuncs/ecs"

type DescribeHaVipsResponse struct {
	ecs.Response
	TotalCount	int  //实例总个数
	PageNumber	int  //实例列表的页码
	PageSize	int  //输入时设置的每页行数
	HaVips	HaVipSetType  //HaVip 详情 HaVipItemType 的集合
}

// TODO API文档里没有
type HaVipSetType struct {

}