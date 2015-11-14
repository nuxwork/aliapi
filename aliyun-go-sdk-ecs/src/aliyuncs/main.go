package main

import(
	"fmt"
	"aliyuncs/ecs"
	"aliyuncs/ecs/model"
)

func main(){
	client   := ecs.NewDefaultClient("h0PSMY8ybRJGPdPe", "KSH5n13i2yIE9RfnZPogcfURrEoRCB", "cn-qingdao")
	
	request  := model.NewDescribeAutoSnapshotPolicyRequest()

	response := client.Do(request)

	if !response.OK() {
		fmt.Println(response.Params())
	}
}