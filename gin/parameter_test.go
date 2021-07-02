// Package ggin
// Created by guoxin in 2020/10/31 10:01 下午
package ggin

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

//func TestGetFieldNames(t *testing.T) {
//    _ = environment.New()
//    d := new(Demo)
//    names, err := GetFieldNames(d)
//    if err != nil {
//        t.Error(err)
//    }
//    t.Log(names)
//}
