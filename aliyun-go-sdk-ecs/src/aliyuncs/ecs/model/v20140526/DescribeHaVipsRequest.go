package v20140526

import "strconv"
import "aliyuncs/core"

type DescribeHaVipsRequest struct {
	core.Request
}

func NewDescribeHaVipsRequest() *DescribeHaVipsRequest {
	r := new(DescribeHaVipsRequest)
	r.Init("Ecs", "2014-05-26", "DescribeHaVips")
	return r
}

func (this *DescribeHaVipsRequest) SetRegionId(v string) {
	this.Set("RegionId", v)
}

func (this *DescribeHaVipsRequest) GetRegionId() string {
	return this.Get("RegionId")
}

func (this *DescribeHaVipsRequest) SetFilternKey(v string, n int) {
	k := "Filter." + strconv.Itoa(n) + ".Key"
	this.Set(k, v)
}

func (this *DescribeHaVipsRequest) GetFilternKey(n int) string {
	k := "Filter." + strconv.Itoa(n) + ".Key"
	return this.Get(k)
}

func (this *DescribeHaVipsRequest) SetFilternValuem(v string, n int, m int) {
	k := "Filter." + strconv.Itoa(n) + ".Value." + strconv.Itoa(m)
	this.Set(k, v)
}

func (this *DescribeHaVipsRequest) GetFilternValuem(n int, m int) string {
	k := "Filter." + strconv.Itoa(n) + ".Value." + strconv.Itoa(m)
	return this.Get(k)
}

func (this *DescribeHaVipsRequest) SetPageNumber(v int) {
	this.SetInt("PageNumber", v)
}

func (this *DescribeHaVipsRequest) GetPageNumber() int {
	return this.GetInt("PageNumber")
}

func (this *DescribeHaVipsRequest) SetPageSize(v int) {
	this.SetInt("PageSize", v)
}

func (this *DescribeHaVipsRequest) GetPageSize() int {
	return this.GetInt("PageSize")
}
