package model

import "aliyuncs/ecs"

type DeleteRouteEntryRequest struct {
	ecs.Request
}

func NewDeleteRouteEntryRequest() *DeleteRouteEntryRequest{
	r := new(DeleteRouteEntryRequest)
	r.Set("Action", "DeleteRouteEntry")
	return r
}

func (this *DeleteRouteEntryRequest) SetRouteTableId(v string){
	this.Set("RouteTableId", v)
}

func (this *DeleteRouteEntryRequest) GetRouteTableId() string {
	return this.Get("RouteTableId")
}

func (this *DeleteRouteEntryRequest) SetDestinationCidrBlock(v string){
	this.Set("DestinationCidrBlock", v)
}

func (this *DeleteRouteEntryRequest) GetDestinationCidrBlock() string {
	return this.Get("DestinationCidrBlock")
}

func (this *DeleteRouteEntryRequest) SetNextHopId(v string){
	this.Set("NextHopId", v)
}

func (this *DeleteRouteEntryRequest) GetNextHopId() string {
	return this.Get("NextHopId")
}
