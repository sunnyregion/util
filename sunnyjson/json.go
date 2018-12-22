package sunnyjson

import "github.com/json-iterator/go"

//这个是为了统一json格式

//Marshal ...
func Marshal(v interface{}) ([]byte, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Marshal(v)

}

//Unmarshal ...
func Unmarshal(data []byte, v interface{}) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Unmarshal(data, v)
}
