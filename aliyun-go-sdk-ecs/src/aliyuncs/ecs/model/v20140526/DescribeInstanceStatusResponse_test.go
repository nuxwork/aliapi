package v20140526

import (
	"testing"
)

func TestDescribeInstanceStatusResponseJson(t *testing.T) {
	data := `
		{
		"Raw": null,
		"HostId": "",
		"Code": "",
		"Message": "",
		"RequestId": "6EF60BEC-0242-43AF-BB20-270359FB54A7",
		"TotalCount": 2,
		"PageNumber": 1,
		"PageSize": 10,
		"InstanceStatuses": {
		    "InstanceStatus": [{
		        "InstanceId": "i-instance1",
		            "Status": "Running"
		        },
		        {
		            "InstanceId": "i-ae4r89pp",
		            "Status": "Stopped"
		        }]
		    }
		}
	`
	var resp DescribeInstanceStatusResponse
	JsonTest(data, &resp, t)
}

func TestDescribeInstanceStatusResponseXml(t *testing.T) {
	data := `
		<DescribeInstanceStatusResponse>
				<HostId></HostId>
				<Code></Code>
				<Message></Message>
		    <RequestId>6EF60BEC-0242-43AF-BB20-270359FB54A7</RequestId>
		    <TotalCount>2</TotalCount>
		    <PageNumber>1</PageNumber>
		    <PageSize>10</PageSize>
		    <InstanceStatuses>
		        <InstanceStatus>
		            <InstanceId>i-instance1</InstanceId>
		                <Status>Running</Status>
		            </InstanceStatus>
		            <InstanceStatus>
		                <InstanceId>i-ae4r89pp</InstanceId>
		                <Status>Stopped</Status>
		        </InstanceStatus>
		    </InstanceStatuses>
		</DescribeInstanceStatusResponse>
	`

	var resp DescribeInstanceStatusResponse
	XmlTest(data, &resp, t)
}
