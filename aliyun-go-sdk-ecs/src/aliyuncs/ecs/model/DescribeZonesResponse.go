package model

import "aliyuncs/ecs"

type DescribeZonesResponse struct {
  ecs.Response
  Zones  ZoneType  //数据中心信息 ZoneType 组成的集合
}

type ZoneSetType struct {
  Zone []ZoneType
}

type ZoneType struct {
  ZoneId  string  //可用区ID
  LocalName  string  //可用区本地语言名
  AvailableResourceCreation  AvailableResourceCreationType  //允许创建的资源类型集合
  AvailableDiskCategories  AvailableDiskCategoriesType  //支持的磁盘种类集合
}

type AvailableResourceCreationType struct {
  ResourceTypes []string
}

type AvailableDiskCategoriesType struct {
  DiskCategories []string
}