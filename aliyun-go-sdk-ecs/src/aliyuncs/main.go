package main

import(
	"fmt"
	"aliyuncs/ecs"
	"aliyuncs/ecs/model"
	"bytes"
)

func main(){
	client   := ecs.NewDefaultClient("h0PSMY8ybRJGPdPe", "KSH5n13i2yIE9RfnZPogcfURrEoRCB", "cn-qingdao")
	
	request  := model.NewDescribeAutoSnapshotPolicyRequest()

	var response model.DescribeAutoSnapshotPolicyResponse

	err := client.Do(request, &response)

	if err != nil && response.OK() {
		fmt.Println(response)
	}

	s := `	

	<nihao 
		<addf
			<ddd
	  `

	fmt.Printf("\n%s\n", escape(s) )
}

func escape(s string) string {
	var buf bytes.Buffer

	for _, r := range s {
		if !shouldEscape(r) {
			buf.WriteRune(r)
		}
	}
	return buf.String()
}

func shouldEscape(r rune) bool {
	if r >=0 && r <= 32 {
		return true
	}

	return false
}