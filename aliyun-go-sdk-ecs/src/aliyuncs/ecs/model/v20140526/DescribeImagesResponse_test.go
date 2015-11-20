package v20140526

import (
	"testing"
)

func TestDescribeImagesResponseJson(t *testing.T) {
	data := `
{
  "Raw": null,
  "HostId": "",
  "Code": "",
  "Message": "",
  "PageNumber": 1, 
  "TotalCount": 39, 
  "PageSize": 10, 
  "RegionId": "cn-qingdao", 
  "RequestId": "FE5823C5-6935-4DF4-8C92-FCE366329627", 
  "Images": {
    "Image": [
      {
        "ImageId": "freebsd1001_64_20G_aliaegis_20150527.vhd", 
        "Description": "freebsd1001_64_20G_aliaegis_20150527.vhd", 
        "ProductCode": "", 
        "OSType": "linux", 
        "Architecture": "x86_64", 
        "OSName": "FreeBSD  10.1 64", 
        "DiskDeviceMappings": {
          "DiskDeviceMapping": [
            {
              "Device": "/dev/xvda", 
              "SnapshotId": "", 
              "Size": "20"
            }
          ]
        }, 
        "ImageOwnerAlias": "system", 
        "Progress": "100%", 
        "Usage": "instance", 
        "CreationTime": "2015-06-19T07:25:42Z", 
        "ImageVersion": "1.0.0", 
        "Status": "Available", 
        "ImageName": "freebsd1001_64_20G_aliaegis_20150527.vhd", 
        "IsCopied": false, 
        "IsSubscribed": false, 
        "Platform": "Freebsd", 
        "Size": 20
      }
    ]
  }
}
	`
	var resp DescribeImagesResponse
	JsonTest(data, &resp, t)
}

func TestDescribeImagesResponseXml(t *testing.T) {
	data := `
<DescribeImagesResponse>
  <HostId></HostId>
  <Code></Code>
  <Message></Message>
  <PageNumber>1</PageNumber>  
  <TotalCount>39</TotalCount>  
  <PageSize>10</PageSize>  
  <RegionId>cn-qingdao</RegionId>  
  <RequestId>1C27A1ED-99AC-49B2-A0F4-D5C7EBFB361A</RequestId>  
  <Images> 
    <Image> 
      <ImageId>freebsd1001_64_20G_aliaegis_20150527.vhd</ImageId>  
      <Description>freebsd1001_64_20G_aliaegis_20150527.vhd</Description>  
      <ProductCode></ProductCode>  
      <OSType>linux</OSType>  
      <Architecture>x86_64</Architecture>  
      <OSName>FreeBSD 10.1 64</OSName>  
      <DiskDeviceMappings> 
        <DiskDeviceMapping> 
          <Device>/dev/xvda</Device>  
          <SnapshotId></SnapshotId>  
          <Size>20</Size> 
        </DiskDeviceMapping> 
      </DiskDeviceMappings>  
      <ImageOwnerAlias>system</ImageOwnerAlias>  
      <Progress>100%</Progress>  
      <Usage>instance</Usage>  
      <CreationTime>2015-06-19T07:25:42Z</CreationTime>  
      <ImageVersion>1.0.0</ImageVersion>  
      <Status>Available</Status>  
      <ImageName>freebsd1001_64_20G_aliaegis_20150527.vhd</ImageName>  
      <IsCopied>false</IsCopied>  
      <IsSubscribed>false</IsSubscribed>  
      <Platform>Freebsd</Platform>  
      <Size>20</Size> 
    </Image>  
  </Images> 
</DescribeImagesResponse>
	`

	var resp DescribeImagesResponse
	XmlTest(data, &resp, t)
}
