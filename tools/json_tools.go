/*
   Created by guoxin in 2020/3/10 4:41 下午
*/
package tools

import "encoding/json"

func ToObject(j string, i interface{}) error {
	e := json.Unmarshal([]byte(j), &i)
	return e
}

func ToJson(i interface{}) (string, error) {
	v, e := json.Marshal(i)
	return string(v), e
}
