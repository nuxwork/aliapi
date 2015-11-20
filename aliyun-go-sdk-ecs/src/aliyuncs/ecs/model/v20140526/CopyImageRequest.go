package v20140526

import "aliyuncs/core"

type CopyImageRequest struct {
	core.Request
}

func NewCopyImageRequest() *CopyImageRequest {
	r := new(CopyImageRequest)
	r.Init("Ecs", "2014-05-26", "CopyImage")
	return r
}

func (this *CopyImageRequest) SetRegionId(v string) {
	this.Set("RegionId", v)
}

func (this *CopyImageRequest) GetRegionId() string {
	return this.Get("RegionId")
}

func (this *CopyImageRequest) SetImageId(v string) {
	this.Set("ImageId", v)
}

func (this *CopyImageRequest) GetImageId() string {
	return this.Get("ImageId")
}

func (this *CopyImageRequest) SetDestinationRegionId(v string) {
	this.Set("DestinationRegionId", v)
}

func (this *CopyImageRequest) GetDestinationRegionId() string {
	return this.Get("DestinationRegionId")
}

func (this *CopyImageRequest) SetDestinationImageName(v string) {
	this.Set("DestinationImageName", v)
}

func (this *CopyImageRequest) GetDestinationImageName() string {
	return this.Get("DestinationImageName")
}

func (this *CopyImageRequest) SetDestinationDescription(v string) {
	this.Set("DestinationDescription", v)
}

func (this *CopyImageRequest) GetDestinationDescription() string {
	return this.Get("DestinationDescription")
}

func (this *CopyImageRequest) SetClientToken(v string) {
	this.Set("ClientToken", v)
}

func (this *CopyImageRequest) GetClientToken() string {
	return this.Get("ClientToken")
}
