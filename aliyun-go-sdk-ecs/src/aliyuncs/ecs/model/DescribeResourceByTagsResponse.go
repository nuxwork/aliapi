package model

import "aliyuncs/ecs"

type DescribeResourceByTagsResponse struct {
  ecs.Response
  PageSize  int  //默认 50，最大 100
  PageNumber  int  //实例状态列表的页码，起始值为 1，默认值为 1
  TotalCount  int  //实例总个数
  Resources  ResourcesSetType  //实际返回内容的结构体列表
}

type ResourcesSetType struct {
	Resource []ResourcesType
}

type ResourcesType struct {
	ResourceId	string  //资源 ID
	ResourceType	string  //资源类型
	RegionId	string  //资源所属地域
}