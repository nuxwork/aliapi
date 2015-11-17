package model

import "aliyuncs/ecs"

type DeleteImageRequest struct {
	ecs.Request
}

func NewDeleteImageRequest() *DeleteImageRequest{
	r := new(DeleteImageRequest)
	r.Set("Action", "DeleteImage")
	return r
}

func (this *DeleteImageRequest) SetRegionId(v string){
	this.Set("RegionId", v)
}

func (this *DeleteImageRequest) GetRegionId() string {
	return this.Get("RegionId")
}

func (this *DeleteImageRequest) SetImageId(v string){
	this.Set("ImageId", v)
}

func (this *DeleteImageRequest) GetImageId() string {
	return this.Get("ImageId")
}
