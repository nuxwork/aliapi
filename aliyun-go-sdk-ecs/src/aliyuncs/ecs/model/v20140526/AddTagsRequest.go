package v20140526

import "strconv"
import "aliyuncs/core"

type AddTagsRequest struct {
	core.Request
}

func NewAddTagsRequest() *AddTagsRequest {
	r := new(AddTagsRequest)
	r.Init("Ecs", "2014-05-26", "AddTags")
	return r
}

func (this *AddTagsRequest) SetResourceId(v string) {
	this.Set("ResourceId", v)
}

func (this *AddTagsRequest) GetResourceId() string {
	return this.Get("ResourceId")
}

func (this *AddTagsRequest) SetResourceType(v string) {
	this.Set("ResourceType", v)
}

func (this *AddTagsRequest) GetResourceType() string {
	return this.Get("ResourceType")
}

func (this *AddTagsRequest) SetRegionId(v string) {
	this.Set("RegionId", v)
}

func (this *AddTagsRequest) GetRegionId() string {
	return this.Get("RegionId")
}

func (this *AddTagsRequest) SetTagnKey(v string, n int) {
	k := "Tag." + strconv.Itoa(n) + ".Key"
	this.Set(k, v)
}

func (this *AddTagsRequest) GetTagnKey(n int) string {
	k := "Tag." + strconv.Itoa(n) + ".Key"
	return this.Get(k)
}

func (this *AddTagsRequest) SetTagnValue(v string, n int) {
	k := "Tag." + strconv.Itoa(n) + ".Value"
	this.Set(k, v)
}

func (this *AddTagsRequest) GetTagnValue(n int) string {
	k := "Tag." + strconv.Itoa(n) + ".Value"
	return this.Get(k)
}
