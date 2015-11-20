package v20140526

import "aliyuncs/core"

type DescribeRegionsResponse struct {
	core.Response
	Regions RegionSetType
}

type RegionSetType struct {
	Region []RegionType
}

type RegionType struct {
	RegionId  string //Region ID
	LocalName string //Region名称
}
