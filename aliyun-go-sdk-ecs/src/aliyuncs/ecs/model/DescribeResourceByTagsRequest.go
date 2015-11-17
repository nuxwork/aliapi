package model

import "strconv"
import "aliyuncs/ecs"

type DescribeResourceByTagsRequest struct {
	ecs.Request
}

func NewDescribeResourceByTagsRequest() *DescribeResourceByTagsRequest{
	r := new(DescribeResourceByTagsRequest)
	r.Set("Action", "DescribeResourceByTags")
	return r
}

func (this *DescribeResourceByTagsRequest) SetResourceType(v string){
	this.Set("ResourceType", v)
}

func (this *DescribeResourceByTagsRequest) GetResourceType() string {
	return this.Get("ResourceType")
}

func (this *DescribeResourceByTagsRequest) SetRegionId(v string){
	this.Set("RegionId", v)
}

func (this *DescribeResourceByTagsRequest) GetRegionId() string {
	return this.Get("RegionId")
}

func (this *DescribeResourceByTagsRequest) SetTagnKey(v string, n int){
	k := "Tag." + strconv.Itoa(n) + ".Key"
	this.Set(k, v)
}

func (this *DescribeResourceByTagsRequest) GetTagnKey(n int) string {
	k := "Tag." + strconv.Itoa(n) + ".Key"
	return this.Get(k)
}

func (this *DescribeResourceByTagsRequest) SetTagnValue(v string, n int){
	k := "Tag." + strconv.Itoa(n) + ".Value"
	this.Set(k, v)
}

func (this *DescribeResourceByTagsRequest) GetTagnValue(n int) string {
	k := "Tag." + strconv.Itoa(n) + ".Value"
	return this.Get(k)
}

func (this *DescribeResourceByTagsRequest) SetPageSize(v int){
	this.SetInt("PageSize", v)
}

func (this *DescribeResourceByTagsRequest) GetPageSize() int {
	return this.GetInt("PageSize")
}

func (this *DescribeResourceByTagsRequest) SetPageNumber(v int){
	this.SetInt("PageNumber", v)
}

func (this *DescribeResourceByTagsRequest) GetPageNumber() int {
	return this.GetInt("PageNumber")
}
