package v20140526

import (
	"testing"
)

func TestDescribeInstancesResponseJson(t *testing.T) {
	data := `
{
	"Raw": null,
	"HostId": "",
	"Code": "",
	"Message": "",
    "PageNumber": 1,
    "TotalCount": 1,
    "PageSize": 10,
    "RequestId": "AF8F5B05-5A16-4715-81EF-42FA1C33B7D3",
    "Instances": {
        "Instance": [
            {
                "ImageId": "ubuntu1404_64_20G_aliaegis_20150325.vhd",
                "InnerIpAddress": {
                    "IpAddress": [
                        "10.163.115.83"
                    ]
                },
                "VlanId": "",
                "InstanceId": "i-28dgx668f",
                "EipAddress": {
                    "IpAddress": "",
                    "AllocationId": "",
                    "InternetChargeType": "",
                    "Bandwidth": 5
                },
                "InternetMaxBandwidthIn": -1,
                "ZoneId": "cn-qingdao-b",
                "InternetChargeType": "PayByTraffic",
                "SerialNumber": "78d04a7d-7f9e-4a64-9f8b-e6e6e1c010a0",
                "IoOptimized": false,
                "InternetMaxBandwidthOut": 1,
                "VpcAttributes": {
                    "NatIpAddress": "",
                    "PrivateIpAddress": {
                        "IpAddress": []
                    },
                    "VSwitchId": "",
                    "VpcId": ""
                },
                "DeviceAvailable": true,
                "SecurityGroupIds": {
                    "SecurityGroupId": [
                        "sg-28oofa14r"
                    ]
                },
                "InstanceName": "ccc",
                "Description": "",
                "InstanceNetworkType": "classic",
                "PublicIpAddress": {
                    "IpAddress": [
                        "114.215.92.209"
                    ]
                },
                "HostName": "iZ28dgx668fZ",
                "InstanceType": "ecs.t1.small",
                "CreationTime": "2015-11-15T13:11Z",
                "Tags": {
                    "Tag": []
                },
                "Status": "Running",
                "ClusterId": "",
                "RegionId": "cn-qingdao",
                "InstanceChargeType": "PostPaid",
                "OperationLocks": {
                    "LockReason": []
                },
                "ExpiredTime": "2999-09-08T16:00Z"
            }
        ]
    }
}
	`
	var resp DescribeInstancesResponse
	JsonTest(data, &resp, t)
}

func TestDescribeInstancesResponseXml(t *testing.T) {
	data := `
<DescribeInstancesResponse> 
  <HostId></HostId>
  <Code></Code>
  <Message></Message>
  <PageNumber>1</PageNumber>  
  <TotalCount>1</TotalCount>  
  <PageSize>10</PageSize>  
  <RequestId>323BF0FF-40DA-4304-8A43-302735A8D8F7</RequestId>  
  <Instances> 
    <Instance> 
      <ImageId>ubuntu1404_64_20G_aliaegis_20150325.vhd</ImageId>  
      <InnerIpAddress> 
        <IpAddress>10.163.115.83</IpAddress> 
      </InnerIpAddress>  
      <VlanId></VlanId>  
      <InstanceId>i-28dgx668f</InstanceId>  
      <EipAddress> 
        <IpAddress></IpAddress>  
        <AllocationId></AllocationId>
        <InternetChargeType></InternetChargeType> 
        <Bandwidth>5</Bandwidth>
      </EipAddress>  
      <InternetMaxBandwidthIn>-1</InternetMaxBandwidthIn>  
      <ZoneId>cn-qingdao-b</ZoneId>  
      <InternetChargeType>PayByTraffic</InternetChargeType>  
      <SerialNumber>78d04a7d-7f9e-4a64-9f8b-e6e6e1c010a0</SerialNumber>  
      <IoOptimized>false</IoOptimized>  
      <InternetMaxBandwidthOut>1</InternetMaxBandwidthOut>  
      <VpcAttributes> 
        <NatIpAddress></NatIpAddress>  
        <PrivateIpAddress></PrivateIpAddress>  
        <VSwitchId></VSwitchId>  
        <VpcId></VpcId> 
      </VpcAttributes>  
      <DeviceAvailable>true</DeviceAvailable>  
      <SecurityGroupIds> 
        <SecurityGroupId>sg-28oofa14r</SecurityGroupId> 
      </SecurityGroupIds>  
      <InstanceName>ccc</InstanceName>  
      <Description></Description> 
      <InstanceNetworkType>classic</InstanceNetworkType>  
      <PublicIpAddress> 
        <IpAddress>114.215.92.209</IpAddress> 
      </PublicIpAddress>  
      <HostName>iZ28dgx668fZ</HostName>  
      <InstanceType>ecs.t1.small</InstanceType>  
      <CreationTime>2015-11-15T13:11Z</CreationTime>  
      <Tags></Tags>  
      <Status>Running</Status>  
      <ClusterId></ClusterId>  
      <RegionId>cn-qingdao</RegionId>  
      <InstanceChargeType>PostPaid</InstanceChargeType>  
      <OperationLocks></OperationLocks>  
      <ExpiredTime>2999-09-08T16:00Z</ExpiredTime> 
    </Instance> 
  </Instances> 
</DescribeInstancesResponse>
	`

	var resp DescribeInstancesResponse
	XmlTest(data, &resp, t)
}
