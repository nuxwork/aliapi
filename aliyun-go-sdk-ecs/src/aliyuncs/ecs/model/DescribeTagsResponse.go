package model

import "aliyuncs/ecs"

type DescribeTagsResponse struct {
	ecs.Response
	PageSize	int  //默认 50，最大 100
	PageNumber	int  //实例状态列表的页码，起始值为 1，默认值为 1
	TotalCount	int  //实例总个数
	Tags	TagItemSetType  //实际返回内容的结构体列表
}

type TagItemSetType struct {
	Tag []TagItemType
}

type TagItemType struct {
	TagKey	string  //Tag 的 Key
	TagValue	string  //Tag 的 Value
}