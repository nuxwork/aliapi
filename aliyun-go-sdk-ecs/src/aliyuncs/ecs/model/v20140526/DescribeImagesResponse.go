package v20140526

import "aliyuncs/core"

type DescribeImagesResponse struct {
	core.Response
	RegionId   string
	TotalCount int
	PageNumber int
	PageSize   int
	Images     ImageTypes
}

type ImageTypes struct {
	Image []ImageType
}

type ImageType struct {
	ImageId            string
	ImageVersion       string
	OSType             string
	Platform           string
	Architecture       string
	ImageName          string
	Description        string
	Size               int
	ImageOwnerAlias    string
	OSName             string
	DiskDeviceMappings DiskDeviceMappingTypes
	ProductCode        string
	IsSubscribed       bool
	Progress           string
	Status             string
	CreationTime       string
	Usage              string
	IsCopied           bool
}

type DiskDeviceMappingTypes struct {
	DiskDeviceMapping []DiskDeviceMappingType
}

type DiskDeviceMappingType struct {
	SnapshotId string
	Size       string
	Device     string
}
