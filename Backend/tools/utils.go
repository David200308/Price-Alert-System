package tools

import "encoding/json"

func StringToJSON(s string) interface{} {
	var res interface{}
	json.Unmarshal([]byte(s), &res)
	return res
}
