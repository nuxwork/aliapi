package model

import "aliyuncs/ecs"

type DescribeSecurityGroupsResponse struct {
	ecs.Response
	TotalCount	int  //	安全组的总数
	PageNumber	int  //	当前页码
	PageSize	int  //	每页行数
	RegionId	string	//安全组所属地域 Id
	SecurityGroups	SecurityGroupItemTypes	//安全组信息，由 SecurityGroupItemType 组成的集合
}

type SecurityGroupItemTypes struct {
	SecurityGroup []SecurityGroupItemType
}

type SecurityGroupItemType struct {
	SecurityGroupId	string  //安全组ID
	SecurityGroupName	string  //安全组名称
	Description	string  //描述信息
	VpcId	string  //安全组所属的专有网络
	CreationTime	string  //创建时间。按照ISO8601标准表示，并需要使用UTC时间。格式为：YYYY-MM-DDThh:mmZ
}