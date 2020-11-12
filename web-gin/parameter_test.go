/*
   Created by guoxin in 2020/10/31 10:09 下午
*/
package web_gin

import (
	"github.com/GuoxinL/gcomponent/environment"
	"testing"
)

type Demo struct {
	Id   int      `query:"id"`
	Name int      `path:"name"`
	Dep  []string `query:"dep"`
	ArrD []Arr
	Person
}
type Arr struct {
}
type Person struct {
	Name string `body:"name"`
	Age  string `body:"age"`
}

func TestGetFieldNames(t *testing.T) {
	new(environment.Configuration).Initialize()
	d := new(Demo)
	names, err := GetFieldNames(d)
	if err != nil {
		t.Error(err)
	}
	t.Log(names)
}
