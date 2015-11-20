package v20140526

import "aliyuncs/core"

type DescribeVSwitchesResponse struct {
	core.Response
	VSwitches  VSwitchType //	VSwitch 详情 VSwitchSetType 组成的集合
	TotalCount int         //列表条条目数
	PageNumber int         //当前页码
	PageSize   int         //当前分页包含多少条目
}

type VSwitchSetType struct {
	VSwitch []VSwitchType
}

type VSwitchType struct {
	VSwitchId               string //交换机ID
	VpcId                   string //交换机所在的专有网络
	Status                  string //交换机状态，包括Pending和Available两种
	CidrBlock               string //交换机的地址
	ZoneId                  string //交换机所在的可用区
	AvailableIpAddressCount int    //	交换机当前可用的IP地址数量
	Description             string //描述，不填则为空，默认值为空，[2,256]英文或中文字符，不能以http:// 和https:// 开头。
	VSwitchName             string //交换机名字，不填则为空，默认值为空，[2,128]英文或中文字符，必须以大小字母或中文开头，可包含数字，”_”或”-”，这个值会展示在控制台。不能以http:// 和https:// 开头。
	CreationTime            string //创建时间。按照ISO8601标准表示，并需要使用UTC时间。格式为：YYYY-MM-DDThh:mmZ
}
