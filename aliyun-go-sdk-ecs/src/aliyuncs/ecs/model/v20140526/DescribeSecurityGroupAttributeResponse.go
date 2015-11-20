package v20140526

import "aliyuncs/core"

type DescribeSecurityGroupAttributeResponse struct {
	core.Response
	SecurityGroupId   string          //	目标安全组 Id
	SecurityGroupName string          //	目标安全组名称
	RegionId          string          //	地域 Id
	Description       string          //	安全组描述信息
	Permissions       PermissionTypes //由 PermissionType 组成的集合，表示安全组中的权限规则
	VpcId             string          //	VpcId，如果有，表示这个安全组是 Vpc 类型的安全组，没有则表示是经典网络里的安全组。
}

type PermissionTypes struct {
	Permission []PermissionType
}

type PermissionType struct {
	IpProtocol              string //	IP协议
	PortRange               string //	端口范围
	SourceCidrIp            string //	源IP地址段，用于Ingress授权
	SourceGroupId           string //	源安全组，用于Ingress授权
	SourceGroupOwnerAccount string //	源安全组所属阿里云账户
	DestCidrIp              string //	目标IP地址段，用于Egress授权
	DestGroupId             string //	目标安全组，用于Egress授权
	DestGroupOwnerAccount   string //	目标安全组所属阿里云账户
	Policy                  string //	授权策略
	NicType                 string //	网络类型
	Priority                string //	规则优先级
	Direction               string //	授权方向
	Description             string //	描述信息
}
