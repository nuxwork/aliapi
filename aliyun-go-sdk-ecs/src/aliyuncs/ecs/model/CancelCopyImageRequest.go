package model

import "aliyuncs/ecs"

type CancelCopyImageRequest struct {
	ecs.Request
}

func NewCancelCopyImageRequest() *CancelCopyImageRequest{
	r := new(CancelCopyImageRequest)
	r.Set("Action", "CancelCopyImage")
	return r
}

func (this *CancelCopyImageRequest) SetRegionId(v string){
	this.Set("RegionId", v)
}

func (this *CancelCopyImageRequest) GetRegionId() string {
	return this.Get("RegionId")
}

func (this *CancelCopyImageRequest) SetImageId(v string){
	this.Set("ImageId", v)
}

func (this *CancelCopyImageRequest) GetImageId() string {
	return this.Get("ImageId")
}
