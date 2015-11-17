package model

import "aliyuncs/ecs"

type CreateImageRequest struct {
	ecs.Request
}

func NewCreateImageRequest() *CreateImageRequest{
	r := new(CreateImageRequest)
	r.Set("Action", "CreateImage")
	return r
}

func (this *CreateImageRequest) SetRegionId(v string){
	this.Set("RegionId", v)
}

func (this *CreateImageRequest) GetRegionId() string {
	return this.Get("RegionId")
}

func (this *CreateImageRequest) SetSnapshotId(v string){
	this.Set("SnapshotId", v)
}

func (this *CreateImageRequest) GetSnapshotId() string {
	return this.Get("SnapshotId")
}

func (this *CreateImageRequest) SetImageName(v string){
	this.Set("ImageName", v)
}

func (this *CreateImageRequest) GetImageName() string {
	return this.Get("ImageName")
}

func (this *CreateImageRequest) SetImageVersion(v string){
	this.Set("ImageVersion", v)
}

func (this *CreateImageRequest) GetImageVersion() string {
	return this.Get("ImageVersion")
}

func (this *CreateImageRequest) SetDescription(v string){
	this.Set("Description", v)
}

func (this *CreateImageRequest) GetDescription() string {
	return this.Get("Description")
}

func (this *CreateImageRequest) SetClientToken(v string){
	this.Set("ClientToken", v)
}

func (this *CreateImageRequest) GetClientToken() string {
	return this.Get("ClientToken")
}
