package lib

import "encoding/json"

func PointerToString(p interface{}) string {
	resp, _ := json.Marshal(p)
	return string(resp)
}