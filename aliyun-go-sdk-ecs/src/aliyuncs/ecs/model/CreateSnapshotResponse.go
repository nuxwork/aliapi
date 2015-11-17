package model

import "aliyuncs/ecs"

type CreateSnapshotResponse struct {
	ecs.Response
	SnapshotId string
}