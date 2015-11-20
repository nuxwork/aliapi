package v20140526

import "aliyuncs/core"

type DescribeImagesRequest struct {
	core.Request
}

func NewDescribeImagesRequest() *DescribeImagesRequest {
	r := new(DescribeImagesRequest)
	r.Init("Ecs", "2014-05-26", "DescribeImages")
	return r
}

func (this *DescribeImagesRequest) SetRegionId(v string) {
	this.Set("RegionId", v)
}

func (this *DescribeImagesRequest) GetRegionId() string {
	return this.Get("RegionId")
}

func (this *DescribeImagesRequest) SetImageId(v string) {
	this.Set("ImageId", v)
}

func (this *DescribeImagesRequest) GetImageId() string {
	return this.Get("ImageId")
}

func (this *DescribeImagesRequest) SetStatus(v string) {
	this.Set("Status", v)
}

func (this *DescribeImagesRequest) GetStatus() string {
	return this.Get("Status")
}

func (this *DescribeImagesRequest) SetSnapshotId(v string) {
	this.Set("SnapshotId", v)
}

func (this *DescribeImagesRequest) GetSnapshotId() string {
	return this.Get("SnapshotId")
}

func (this *DescribeImagesRequest) SetImageName(v string) {
	this.Set("ImageName", v)
}

func (this *DescribeImagesRequest) GetImageName() string {
	return this.Get("ImageName")
}

func (this *DescribeImagesRequest) SetImageOwnerAlias(v string) {
	this.Set("ImageOwnerAlias", v)
}

func (this *DescribeImagesRequest) GetImageOwnerAlias() string {
	return this.Get("ImageOwnerAlias")
}

func (this *DescribeImagesRequest) SetUsage(v string) {
	this.Set("Usage", v)
}

func (this *DescribeImagesRequest) GetUsage() string {
	return this.Get("Usage")
}

func (this *DescribeImagesRequest) SetPageNumber(v string) {
	this.Set("PageNumber", v)
}

func (this *DescribeImagesRequest) GetPageNumber() string {
	return this.Get("PageNumber")
}

func (this *DescribeImagesRequest) SetPageSize(v string) {
	this.Set("PageSize", v)
}

func (this *DescribeImagesRequest) GetPageSize() string {
	return this.Get("PageSize")
}
