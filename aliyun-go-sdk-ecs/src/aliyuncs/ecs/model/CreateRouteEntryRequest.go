package model

import "aliyuncs/ecs"

type CreateRouteEntryRequest struct {
	ecs.Request
}

func NewCreateRouteEntryRequest() *CreateRouteEntryRequest{
	r := new(CreateRouteEntryRequest)
	r.Set("Action", "CreateRouteEntry")
	return r
}

func (this *CreateRouteEntryRequest) SetRouteTableId(v string){
	this.Set("RouteTableId", v)
}

func (this *CreateRouteEntryRequest) GetRouteTableId() string {
	return this.Get("RouteTableId")
}

func (this *CreateRouteEntryRequest) SetDestinationCidrBlock(v string){
	this.Set("DestinationCidrBlock", v)
}

func (this *CreateRouteEntryRequest) GetDestinationCidrBlock() string {
	return this.Get("DestinationCidrBlock")
}

func (this *CreateRouteEntryRequest) SetNextHopType(v string){
	this.Set("NextHopType", v)
}

func (this *CreateRouteEntryRequest) GetNextHopType() string {
	return this.Get("NextHopType")
}

func (this *CreateRouteEntryRequest) SetNextHopId(v string){
	this.Set("NextHopId", v)
}

func (this *CreateRouteEntryRequest) GetNextHopId() string {
	return this.Get("NextHopId")
}

func (this *CreateRouteEntryRequest) SetClientToken(v string){
	this.Set("ClientToken", v)
}

func (this *CreateRouteEntryRequest) GetClientToken() string {
	return this.Get("ClientToken")
}
