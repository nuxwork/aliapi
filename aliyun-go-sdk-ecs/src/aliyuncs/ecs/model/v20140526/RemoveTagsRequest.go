package v20140526

import "strconv"
import "aliyuncs/core"

type RemoveTagsRequest struct {
	core.Request
}

func NewRemoveTagsRequest() *RemoveTagsRequest {
	r := new(RemoveTagsRequest)
	r.Init("Ecs", "2014-05-26", "RemoveTags")
	return r
}

func (this *RemoveTagsRequest) SetResourceId(v string) {
	this.Set("ResourceId", v)
}

func (this *RemoveTagsRequest) GetResourceId() string {
	return this.Get("ResourceId")
}

func (this *RemoveTagsRequest) SetResourceType(v string) {
	this.Set("ResourceType", v)
}

func (this *RemoveTagsRequest) GetResourceType() string {
	return this.Get("ResourceType")
}

func (this *RemoveTagsRequest) SetRegionId(v string) {
	this.Set("RegionId", v)
}

func (this *RemoveTagsRequest) GetRegionId() string {
	return this.Get("RegionId")
}

func (this *RemoveTagsRequest) SetTagnKey(v string, n int) {
	k := "Tag." + strconv.Itoa(n) + ".Key"
	this.Set(k, v)
}

func (this *RemoveTagsRequest) GetTagnKey(n int) string {
	k := "Tag." + strconv.Itoa(n) + ".Key"
	return this.Get(k)
}

func (this *RemoveTagsRequest) SetTagnValue(v string, n int) {
	k := "Tag." + strconv.Itoa(n) + ".Value"
	this.Set(k, v)
}

func (this *RemoveTagsRequest) GetTagnValue(n int) string {
	k := "Tag." + strconv.Itoa(n) + ".Value"
	return this.Get(k)
}
