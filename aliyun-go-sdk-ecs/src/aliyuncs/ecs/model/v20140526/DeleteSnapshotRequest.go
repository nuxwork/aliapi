package v20140526

import "aliyuncs/core"

type DeleteSnapshotRequest struct {
	core.Request
}

func NewDeleteSnapshotRequest() *DeleteSnapshotRequest {
	r := new(DeleteSnapshotRequest)
	r.Init("Ecs", "2014-05-26", "DeleteSnapshot")
	return r
}

func (this *DeleteSnapshotRequest) SetSnapshotId(v string) {
	this.Set("SnapshotId", v)
}

func (this *DeleteSnapshotRequest) GetSnapshotId() string {
	return this.Get("SnapshotId")
}
