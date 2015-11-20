package v20140526

import "aliyuncs/core"

type CancelCopyImageRequest struct {
	core.Request
}

func NewCancelCopyImageRequest() *CancelCopyImageRequest {
	r := new(CancelCopyImageRequest)
	r.Init("Ecs", "2014-05-26", "CancelCopyImage")
	return r
}

func (this *CancelCopyImageRequest) SetRegionId(v string) {
	this.Set("RegionId", v)
}

func (this *CancelCopyImageRequest) GetRegionId() string {
	return this.Get("RegionId")
}

func (this *CancelCopyImageRequest) SetImageId(v string) {
	this.Set("ImageId", v)
}

func (this *CancelCopyImageRequest) GetImageId() string {
	return this.Get("ImageId")
}
