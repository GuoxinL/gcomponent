/*
   Environment test
   Created by guoxin in 2020/11/10 8:34 上午
*/
package environment

import (
    "encoding/json"
    "fmt"
    "testing"
)

// Environment variables FOO_COO=321
func TestInit(t *testing.T) {
    application := newApplicationFile("", "")

    ins := struct {
        Foo struct {
            Boo string   ` mapstructure:"boo" default:""`
            Coo int      ` mapstructure:"coo" default:""`
            Doo []string `mapstructure:"doo"`
        } `mapstructure:"foo"`
    }{}
    application.SetDefault("foo.coo", 456)
    application.AutomaticEnv()
    err := application.Unmarshal(&ins)
    coo := application.GetString("foo.coo")
    fmt.Println(coo)
    if err != nil {
        t.Error("UnmarshalKey error")
    }
    j, err := json.Marshal(ins)
    if err != nil {
        t.Error("ToJson error")
    }
    t.Log("yaml print: ", string(j))
}

//func ParseStruct(t reflect.Type, tag string) {
//    if t.Kind() == reflect.Ptr {
//        t = t.Elem()
//    }
//    if t.Kind() != reflect.Struct {
//        return
//    }
//
//    for i := 0; i < t.NumField(); i++ {
//        f := t.Field(i)
//        value := f.Tag.Get(tag)
//        ft := t.Type
//        if ft.Kind() == reflect.Ptr {
//            ft = ft.Elem()
//        }
//
//        // It seems that you don't want a tag from an struct field; only other fields' tags are needed
//        if ft.Kind() != reflect.Struct {
//            if len(value) != 0 {
//                fmt.Println(value)
//            }
//            continue
//        }
//
//        ParseStruct(ft, tag)
//    }
//}
