package v20140526

import (
	"testing"
)

func TestDescribeImageSharePermissionResponseJson(t *testing.T) {
	data := `
{
	"Raw": null,
	"HostId": "",
	"Code": "",
	"Message": "",
  "Accounts": {
    "Account": [
      {
        "AliyunId": "1265133900679922"
      }, 
      {
        "AliyunId": "1796187892626608"
      }
    ]
  }, 
  "ImageId": "m-2848re4jk", 
  "PageNumber": 1, 
  "TotalCount": 2, 
  "PageSize": 10, 
  "RegionId": "cn-qingdao", 
  "RequestId": "F38DE477-1C81-4109-B873-5617C5B7FDD8", 
  "ShareGroups": {
    "ShareGroup": [ ]
  }
}
	`
	var resp DescribeImageSharePermissionResponse
	JsonTest(data, &resp, t)
}

func TestDescribeImageSharePermissionResponseXml(t *testing.T) {
	data := `
<DescribeImageSharePermissionResponse>
	<HostId></HostId>
	<Code></Code>
	<Message></Message> 
  <Accounts> 
    <Account> 
      <AliyunId>1265133900679922</AliyunId> 
    </Account>  
    <Account> 
      <AliyunId>1796187892626608</AliyunId> 
    </Account> 
  </Accounts>  
  <ImageId>m-2848re4jk</ImageId>  
  <PageNumber>1</PageNumber>  
  <TotalCount>2</TotalCount>  
  <PageSize>10</PageSize>  
  <RegionId>cn-qingdao</RegionId>  
  <RequestId>3C5D9C9D-A11A-401A-8C7D-5EA96354A07E</RequestId>  
  <ShareGroups></ShareGroups>
</DescribeImageSharePermissionResponse>
	`

	var resp DescribeImageSharePermissionResponse
	XmlTest(data, &resp, t)
}
