package v20140526

import "aliyuncs/core"

type CreateVSwitchRequest struct {
	core.Request
}

func NewCreateVSwitchRequest() *CreateVSwitchRequest {
	r := new(CreateVSwitchRequest)
	r.Init("Ecs", "2014-05-26", "CreateVSwitch")
	return r
}

func (this *CreateVSwitchRequest) SetZoneId(v string) {
	this.Set("ZoneId", v)
}

func (this *CreateVSwitchRequest) GetZoneId() string {
	return this.Get("ZoneId")
}

func (this *CreateVSwitchRequest) SetCidrBlock(v string) {
	this.Set("CidrBlock", v)
}

func (this *CreateVSwitchRequest) GetCidrBlock() string {
	return this.Get("CidrBlock")
}

func (this *CreateVSwitchRequest) SetVpcId(v string) {
	this.Set("VpcId", v)
}

func (this *CreateVSwitchRequest) GetVpcId() string {
	return this.Get("VpcId")
}

func (this *CreateVSwitchRequest) SetVSwitchName(v string) {
	this.Set("VSwitchName", v)
}

func (this *CreateVSwitchRequest) GetVSwitchName() string {
	return this.Get("VSwitchName")
}

func (this *CreateVSwitchRequest) SetDescription(v string) {
	this.Set("Description", v)
}

func (this *CreateVSwitchRequest) GetDescription() string {
	return this.Get("Description")
}

func (this *CreateVSwitchRequest) SetClientToken(v string) {
	this.Set("ClientToken", v)
}

func (this *CreateVSwitchRequest) GetClientToken() string {
	return this.Get("ClientToken")
}
