package BlockChain

import (
	"encoding/json"
)

func ToString(s interface{}) string  {
	bytes, e := json.Marshal(s)
	if e != nil {
		panic("序列化出错拉"+e.Error())
	}
	result :=string(bytes)
	//fmt.Println(result)
	return result
}