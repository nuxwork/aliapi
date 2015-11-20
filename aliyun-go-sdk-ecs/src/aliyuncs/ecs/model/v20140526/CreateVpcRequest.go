package v20140526

import "aliyuncs/core"

type CreateVpcRequest struct {
	core.Request
}

func NewCreateVpcRequest() *CreateVpcRequest {
	r := new(CreateVpcRequest)
	r.Init("Ecs", "2014-05-26", "CreateVpc")
	return r
}

func (this *CreateVpcRequest) SetRegionId(v string) {
	this.Set("RegionId", v)
}

func (this *CreateVpcRequest) GetRegionId() string {
	return this.Get("RegionId")
}

func (this *CreateVpcRequest) SetCidrBlock(v string) {
	this.Set("CidrBlock", v)
}

func (this *CreateVpcRequest) GetCidrBlock() string {
	return this.Get("CidrBlock")
}

func (this *CreateVpcRequest) SetVpcName(v string) {
	this.Set("VpcName", v)
}

func (this *CreateVpcRequest) GetVpcName() string {
	return this.Get("VpcName")
}

func (this *CreateVpcRequest) SetDescription(v string) {
	this.Set("Description", v)
}

func (this *CreateVpcRequest) GetDescription() string {
	return this.Get("Description")
}

func (this *CreateVpcRequest) SetClientToken(v string) {
	this.Set("ClientToken", v)
}

func (this *CreateVpcRequest) GetClientToken() string {
	return this.Get("ClientToken")
}
