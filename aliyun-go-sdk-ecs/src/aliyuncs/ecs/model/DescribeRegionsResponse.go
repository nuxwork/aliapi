package model

import "aliyuncs/ecs"

type DescribeRegionsResponse struct {
  ecs.Response
  Regions RegionSetType
}

type RegionSetType struct {
	Region []RegionType
}

type RegionType struct {
  RegionId  string  //Region ID
  LocalName  string  //Region名称
}