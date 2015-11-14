package model

import "aliyuncs/ecs"

type DeleteSnapshotRequest struct {
	ecs.Request
}

func NewDeleteSnapshotRequest() *DeleteSnapshotRequest{
	r := new(DeleteSnapshotRequest)
	r.Set("Action", "DeleteSnapshot")
	return r
}

func (this *DeleteSnapshotRequest) SetSnapshotId(v string){
	this.Set("SnapshotId", v)
}

func (this *DeleteSnapshotRequest) GetSnapshotId() string {
	return this.Get("SnapshotId")
}
