package v20140526

import "aliyuncs/core"

type ModifyInstanceNetworkSpecRequest struct {
	core.Request
}

func NewModifyInstanceNetworkSpecRequest() *ModifyInstanceNetworkSpecRequest {
	r := new(ModifyInstanceNetworkSpecRequest)
	r.Init("Ecs", "2014-05-26", "ModifyInstanceNetworkSpec")
	return r
}

func (this *ModifyInstanceNetworkSpecRequest) SetInstanceId(v string) {
	this.Set("InstanceId", v)
}

func (this *ModifyInstanceNetworkSpecRequest) GetInstanceId() string {
	return this.Get("InstanceId")
}

func (this *ModifyInstanceNetworkSpecRequest) SetInternetMaxBandwidthOut(v int) {
	this.SetInt("InternetMaxBandwidthOut", v)
}

func (this *ModifyInstanceNetworkSpecRequest) GetInternetMaxBandwidthOut() int {
	return this.GetInt("InternetMaxBandwidthOut")
}

func (this *ModifyInstanceNetworkSpecRequest) SetInternetMaxBandwidthIn(v int) {
	this.SetInt("InternetMaxBandwidthIn", v)
}

func (this *ModifyInstanceNetworkSpecRequest) GetInternetMaxBandwidthIn() int {
	return this.GetInt("InternetMaxBandwidthIn")
}
