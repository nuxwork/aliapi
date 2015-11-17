package model

import (
	"testing"
)

func TestDescribeSnapshotsResponseJson(t *testing.T) {
	data := `
{
	"Raw": null,
	"HostId": "",
	"Code": "",
	"Message": "",
  "PageNumber": 1,
  "PageSize": 2,
  "RegionId": "cn-aingdao",
  "RequestId": "659F91C6-1949-43B0-90C4-B6342CA757D5",
  "Snapshots": {
    "Snapshot": [
      {
        "CreationTime": "2015-07-30T05:00:14Z",
        "Progress": "100%",
        "SnapshotId": "s-943ypfgic",
        "SnapshotName": "auto_20150730_3",
        "SourceDiskId": "d-944qyqjfa",
        "SourceDiskSize": 20,
        "SourceDiskType": "system",
        "Status": "accomplished",
        "Usage": "none",
        "Description":"",
        "ProductCode": ""
      },
      {
        "CreationTime": "2015-07-30T05:00:14Z",
        "Progress": "100%",
        "SnapshotId": "s-94osg320e",
        "SnapshotName": "auto_20150730_3",
        "SourceDiskId": "d-94j355jsq",
        "SourceDiskSize": 20,
        "SourceDiskType": "system",
        "Status": "accomplished",
        "Usage": "none",
        "Description":"",
        "ProductCode": ""
      }
    ]
  },
  "TotalCount": 36
}
	`
	var resp DescribeSnapshotsResponse
	JsonTest(data, &resp, t)
}

func TestDescribeSnapshotsResponseXml(t *testing.T) {
	data := `
<DescribeSnapshotsResponse>
		<HostId></HostId>
		<Code></Code>
		<Message></Message>
		<RegionId></RegionId>
    <Snapshots>
        <Snapshot>
            <CreationTime>2014-07-24T13:00:52Z</CreationTime>
            <Description></Description>
            <SourceDiskId>d-23x0r79qy</SourceDiskId>
            <SourceDiskType>DATA</SourceDiskType>
            <ProductCode></ProductCode>
            <SnapshotName>auto_20140724_2</SnapshotName>
            <Progress>100%</Progress>
            <SourceDiskSize>50</SourceDiskSize>
            <Status>accomplished</Status>
            <SnapshotId>s-23f2i9s4t</SnapshotId>
            <Usage>none</Usage>
        </Snapshot>
        <Snapshot>
            <CreationTime>2014-07-24T13:00:42Z</CreationTime>
            <Description></Description>
            <SourceDiskId>101-70105379</SourceDiskId>
            <SourceDiskType>DATA</SourceDiskType>
            <ProductCode></ProductCode>
            <SnapshotName>auto_20140724_2</SnapshotName>
            <Progress>100%</Progress>
            <SourceDiskSize>5</SourceDiskSize>
            <Status>accomplished</Status>
            <SnapshotId>s-23izto5qm</SnapshotId>
            <Usage>none</Usage>
        </Snapshot>
    </Snapshots>
    <PageNumber>1</PageNumber>
    <PageSize>2</PageSize>
    <TotalCount>36</TotalCount>
    <RequestId>2F409D67-329A-4405-B924-2FC28566B366</RequestId>
</DescribeSnapshotsResponse>
	`

	var resp DescribeSnapshotsResponse
	XmlTest(data, &resp, t)
}