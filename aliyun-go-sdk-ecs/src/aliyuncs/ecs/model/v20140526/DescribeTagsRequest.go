package v20140526

import "strconv"
import "aliyuncs/core"

type DescribeTagsRequest struct {
	core.Request
}

func NewDescribeTagsRequest() *DescribeTagsRequest {
	r := new(DescribeTagsRequest)
	r.Init("Ecs", "2014-05-26", "DescribeTags")
	return r
}

func (this *DescribeTagsRequest) SetResourceId(v string) {
	this.Set("ResourceId", v)
}

func (this *DescribeTagsRequest) GetResourceId() string {
	return this.Get("ResourceId")
}

func (this *DescribeTagsRequest) SetResourceType(v string) {
	this.Set("ResourceType", v)
}

func (this *DescribeTagsRequest) GetResourceType() string {
	return this.Get("ResourceType")
}

func (this *DescribeTagsRequest) SetRegionId(v string) {
	this.Set("RegionId", v)
}

func (this *DescribeTagsRequest) GetRegionId() string {
	return this.Get("RegionId")
}

func (this *DescribeTagsRequest) SetTagnKey(v string, n int) {
	k := "Tag." + strconv.Itoa(n) + ".Key"
	this.Set(k, v)
}

func (this *DescribeTagsRequest) GetTagnKey(n int) string {
	k := "Tag." + strconv.Itoa(n) + ".Key"
	return this.Get(k)
}

func (this *DescribeTagsRequest) SetTagnValue(v string, n int) {
	k := "Tag." + strconv.Itoa(n) + ".Value"
	this.Set(k, v)
}

func (this *DescribeTagsRequest) GetTagnValue(n int) string {
	k := "Tag." + strconv.Itoa(n) + ".Value"
	return this.Get(k)
}

func (this *DescribeTagsRequest) SetPageSize(v int) {
	this.SetInt("PageSize", v)
}

func (this *DescribeTagsRequest) GetPageSize() int {
	return this.GetInt("PageSize")
}

func (this *DescribeTagsRequest) SetPageNumber(v int) {
	this.SetInt("PageNumber", v)
}

func (this *DescribeTagsRequest) GetPageNumber() int {
	return this.GetInt("PageNumber")
}
