package v20140526

import "aliyuncs/core"

type AllocateEipAddressResponse struct {
	core.Response
	EipAddress   string
	AllocationId string
}
