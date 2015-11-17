package model

import "aliyuncs/ecs"

type DescribeImageSharePermissionRequest struct {
	ecs.Request
}

func NewDescribeImageSharePermissionRequest() *DescribeImageSharePermissionRequest{
	r := new(DescribeImageSharePermissionRequest)
	r.Set("Action", "DescribeImageSharePermission")
	return r
}

func (this *DescribeImageSharePermissionRequest) SetRegionId(v string){
	this.Set("RegionId", v)
}

func (this *DescribeImageSharePermissionRequest) GetRegionId() string {
	return this.Get("RegionId")
}

func (this *DescribeImageSharePermissionRequest) SetImageId(v string){
	this.Set("ImageId", v)
}

func (this *DescribeImageSharePermissionRequest) GetImageId() string {
	return this.Get("ImageId")
}

func (this *DescribeImageSharePermissionRequest) SetPageNumber(v int){
	this.SetInt("PageNumber", v)
}

func (this *DescribeImageSharePermissionRequest) GetPageNumber() int {
	return this.GetInt("PageNumber")
}

func (this *DescribeImageSharePermissionRequest) SetPageSize(v int){
	this.SetInt("PageSize", v)
}

func (this *DescribeImageSharePermissionRequest) GetPageSize() int {
	return this.GetInt("PageSize")
}
