package v20140526

import "aliyuncs/core"

type DeleteImageRequest struct {
	core.Request
}

func NewDeleteImageRequest() *DeleteImageRequest {
	r := new(DeleteImageRequest)
	r.Init("Ecs", "2014-05-26", "DeleteImage")
	return r
}

func (this *DeleteImageRequest) SetRegionId(v string) {
	this.Set("RegionId", v)
}

func (this *DeleteImageRequest) GetRegionId() string {
	return this.Get("RegionId")
}

func (this *DeleteImageRequest) SetImageId(v string) {
	this.Set("ImageId", v)
}

func (this *DeleteImageRequest) GetImageId() string {
	return this.Get("ImageId")
}
