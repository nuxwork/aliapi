package main

import (
	acs "aliyuncs/core"
	ecs "aliyuncs/ecs/model/v20140526"
	"fmt"
)

func main() {
	profile := acs.NewProfile("h0PSMY8ybRJGPdPe", "KSH5n13i2yIE9RfnZPogcfURrEoRCB", "cn-qingdao")
	profile.SetFormat(acs.FORMAT_JSON)

	client := acs.NewClient(profile)

	request := ecs.NewStartInstanceRequest()
	// request.SetImageId("m-2848re4jk")
	// request.SetRegionId("cn-qingdao")

	var response ecs.StartInstanceResponse

	err := client.Do(request, &response)

	fmt.Println(err)
	fmt.Println()
	fmt.Println(response)

	if err == nil && response.OK() {
		fmt.Println("OK!!!")
	}else{
		fmt.Printf("\nCode: %s\nMessage: %s\n", response.Code, response.Message)
	}
}
