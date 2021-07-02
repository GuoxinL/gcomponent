// Package ggin
// Created by guoxin in 2020/10/31 10:01 下午
package ggin

import (
	"github.com/gin-gonic/gin"
)

// ResultBuilder
// request
//	1. query
//	2. body
//	3. path
//	4. header
//
// response
//	1. code
//	2. body
type ResultBuilder interface {
}

type Parameter struct {
	*gin.Context
}

//// GetFieldNames 获取结构体中字段的名称
//func GetFieldNames(i interface{}) (map[string]string, error) {
//    t := reflect.TypeOf(i)
//    t = realType(t)
//    if t.Kind() != reflect.Struct {
//        return nil, errors.New("Check type error not Struct")
//    }
//    num := t.NumField()
//    m := make(map[string]string, num)
//    for i := 0; i < num; i++ {
//        //m[i] = t.Field(i).Name
//        if t.Field(i).Tag == "" {
//            of := reflect.TypeOf(t.Field(i).Type)
//            of = realType(of)
//            if of.Kind() == reflect.Struct {
//                fmt.Println(of)
//            }
//        } else {
//            //tag := t.Field(i).Tag
//            //
//            //lookup, ok := tag.Lookup("query")
//            //if ok {
//            //	m
//            //	[]
//            //}
//        }
//
//        fmt.Println(i, "Name", t.Field(i).Name)
//        fmt.Println(i, "Index", t.Field(i).Index)
//        fmt.Println(i, "Type", t.Field(i).Type)
//        fmt.Println(i, "Anonymous", t.Field(i).Anonymous)
//        fmt.Println(i, "Offset", t.Field(i).Offset)
//        fmt.Println(i, "PkgPath", t.Field(i).PkgPath)
//        fmt.Println(i, "Tag", t.Field(i).Tag)
//    }
//    return m, nil
//}
//
//func realType(t reflect.Type) reflect.Type {
//    if t.Kind() == reflect.Ptr {
//        t = t.Elem()
//        return realType(t)
//    } else {
//        return t
//    }
//}
