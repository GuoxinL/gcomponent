/*
   Created by guoxin in 2020/1/10 11:16 上午
*/
package tools

import "sync"

type ConcurrentMap struct {
	d map[string]interface{}
	l sync.Mutex
}

//func (this *ConcurrentMap) put(key string, value interface{}) {
//	this.d[key] = value
//}
//func (this *ConcurrentMap) get(key string) interface{} {
//	value, ok := this.d[key]
//	if ok {
//		return value
//	}
//	return nil
//}

func (this *ConcurrentMap) Put(key string, value interface{}) {
	this.l.Lock()
	defer this.l.Unlock()
	this.d[key] = value
}

func (this *ConcurrentMap) Get(key string) interface{} {
	this.l.Lock()
	defer this.l.Unlock()
	value := this.d[key]
	return value
}

func (this *ConcurrentMap) Rm(key string) {
	this.l.Lock()
	defer this.l.Unlock()
	delete(this.d, key)

}

func (this *ConcurrentMap) Size() int {
	this.l.Lock()
	defer this.l.Unlock()
	return len(this.d)
}

func (this *ConcurrentMap) Merge(source map[string]interface{}) {
	this.l.Lock()
	defer this.l.Unlock()
	this.d = CopyMap(source)
}

func (this *ConcurrentMap) Copy() *ConcurrentMap {
	newMap := NewConcurrentMap()
	for k, v := range this.d {
		newMap.d[k] = v
	}
	return newMap
}

/**
从ConcurrentMap中获得线程不安全map
*/
func (this *ConcurrentMap) GetMap() map[string]interface{} {
	this.l.Lock()
	defer this.l.Unlock()
	maps := make(map[string]interface{})
	for k, v := range this.d {
		maps[k] = v
	}
	return maps
}

/**
从ConcurrentMap中获得线程不安全map
*/
func (this *ConcurrentMap) GetAndRemoveMap() map[string]interface{} {
	this.l.Lock()
	defer this.l.Unlock()
	maps := make(map[string]interface{})
	for k, v := range this.d {
		maps[k] = v
		delete(this.d, k)
	}
	return maps
}

func CopyMap(source map[string]interface{}) map[string]interface{} {
	maps := make(map[string]interface{})
	for k, v := range source {
		maps[k] = v
	}
	return maps
}

func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		d: make(map[string]interface{}),
		l: sync.Mutex{},
	}
}
func NewConcurrentMapData(source map[string]interface{}) *ConcurrentMap {
	return &ConcurrentMap{
		d: CopyMap(source),
		l: sync.Mutex{},
	}
}

/**
key: interface
value: interface
*/
type ConcurrentMap0 struct {
	d map[interface{}]interface{}
	l sync.Mutex
}

func (this *ConcurrentMap0) Put(key interface{}, value interface{}) {
	this.l.Lock()
	defer this.l.Unlock()
	this.d[key] = value
}

func (this *ConcurrentMap0) Get(key interface{}) interface{} {
	this.l.Lock()
	defer this.l.Unlock()
	value := this.d[key]
	return value
}

func (this *ConcurrentMap0) Size() int {
	this.l.Lock()
	defer this.l.Unlock()
	return len(this.d)
}

func (this *ConcurrentMap0) Merge(source map[interface{}]interface{}) {
	this.l.Lock()
	defer this.l.Unlock()
	this.d = CopyMap0(source)
}

func (this *ConcurrentMap0) Copy() *ConcurrentMap0 {
	this.l.Lock()
	defer this.l.Unlock()
	newMap := NewConcurrentMap0()
	for k, v := range this.d {
		newMap.d[k] = v
	}
	return newMap
}

func CopyMap0(source map[interface{}]interface{}) map[interface{}]interface{} {
	maps := make(map[interface{}]interface{})
	for k, v := range source {
		maps[k] = v
	}
	return maps
}
func NewConcurrentMap0() *ConcurrentMap0 {
	return &ConcurrentMap0{
		d: make(map[interface{}]interface{}),
		l: sync.Mutex{},
	}
}
func NewConcurrentMapData0(source map[interface{}]interface{}) *ConcurrentMap0 {
	return &ConcurrentMap0{
		d: CopyMap0(source),
		l: sync.Mutex{},
	}
}
