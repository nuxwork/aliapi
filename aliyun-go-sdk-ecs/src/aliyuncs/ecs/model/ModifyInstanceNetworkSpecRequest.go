package model

import "aliyuncs/ecs"

type ModifyInstanceNetworkSpecRequest struct {
	ecs.Request
}

func NewModifyInstanceNetworkSpecRequest() *ModifyInstanceNetworkSpecRequest{
	r := new(ModifyInstanceNetworkSpecRequest)
	r.Set("Action", "ModifyInstanceNetworkSpec")
	return r
}

func (this *ModifyInstanceNetworkSpecRequest) SetInstanceId(v string){
	this.Set("InstanceId", v)
}

func (this *ModifyInstanceNetworkSpecRequest) GetInstanceId() string {
	return this.Get("InstanceId")
}

func (this *ModifyInstanceNetworkSpecRequest) SetInternetMaxBandwidthOut(v int){
	this.SetInt("InternetMaxBandwidthOut", v)
}

func (this *ModifyInstanceNetworkSpecRequest) GetInternetMaxBandwidthOut() int {
	return this.GetInt("InternetMaxBandwidthOut")
}

func (this *ModifyInstanceNetworkSpecRequest) SetInternetMaxBandwidthIn(v int){
	this.SetInt("InternetMaxBandwidthIn", v)
}

func (this *ModifyInstanceNetworkSpecRequest) GetInternetMaxBandwidthIn() int {
	return this.GetInt("InternetMaxBandwidthIn")
}
