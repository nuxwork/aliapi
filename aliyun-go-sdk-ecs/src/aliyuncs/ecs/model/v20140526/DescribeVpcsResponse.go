package v20140526

import "aliyuncs/core"

type DescribeVpcsResponse struct {
	core.Response
	Vpcs       VpcSetType //	Vpc 详情 VpcSetType 组成的集合
	TotalCount string     //列表条条目数
	PageNumber int        //当前页码
	PageSize   int        //当前分页包含多少条目
}

type VpcSetType struct {
	Vpc []VpcType
}

type VpcType struct {
	VpcId        string //VpcId
	RegionId     string //VPC所在的地域
	Status       string //VPC状态，包括Pending和Available两种
	VpcName      string //VPC名称，不填则为空，默认值为空，[2,128]英文或中文字符，必须以大小字母或中文开头，可包含数字，”_”或”-”，这个值会展示在控制台。不能以http:// 和https:// 开头。
	VSwitchIds   string //VSwitchId列表
	CidrBlock    string //VPC的网段地址
	VRouterId    string //VRouter的Id
	Description  string //描述，不填则为空，默认值为空，[2,256]英文或中文字符，不能以http:// 和https:// 开头。
	CreationTime string //创建时间。按照ISO8601标准表示，并需要使用UTC时间。格式为：YYYY-MM-DDThh:mmZ
}
