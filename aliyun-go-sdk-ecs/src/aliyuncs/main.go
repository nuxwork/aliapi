package main

import(
	"fmt"
	"aliyuncs/ecs"
	"aliyuncs/ecs/model"
)

func main(){
	profile := ecs.NewDefaultProfile("h0PSMY8ybRJGPdPe", "KSH5n13i2yIE9RfnZPogcfURrEoRCB", "cn-qingdao")
	// profile.SetFormat(ecs.FORMAT_XML)
	profile.SetFormat(ecs.FORMAT_JSON)

	client  := ecs.NewClient(profile)
	
	request := model.NewDescribeImageSharePermissionRequest()
	request.SetImageId("m-2848re4jk")

	request.SetRegionId("cn-qingdao")

	var response model.DescribeImageSharePermissionResponse

	err := client.Do(request, &response)

	if err != nil && response.OK() {
		fmt.Println(response)
	}
}
