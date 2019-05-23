package main

import (
	"fmt"
	"reflect"
	"sync"
)

type ConcurrentMap struct {
	m         sync.Map
	keyType   reflect.Type
	valueType reflect.Type
}

func (cMap *ConcurrentMap) Load(key interface{}) (value interface{}, ok bool) {
	if reflect.TypeOf(key) != cMap.keyType {
		return // nil, false
	}

	return cMap.m.Load(key)
}

func (cMap *ConcurrentMap) Store(key, value interface{}) {
	if reflect.TypeOf(key) != cMap.keyType {
		panic(fmt.Errorf("wrong key type: %v", reflect.TypeOf(key)))
	}

	if reflect.TypeOf(value) != cMap.valueType {
		panic(fmt.Errorf("wrong value type: %v", reflect.TypeOf(value)))
	}
	cMap.m.Store(key, value)
}

func main() {
	m := ConcurrentMap{
		keyType:   reflect.TypeOf(1),
		valueType: reflect.TypeOf("1"),
	}
	m.Store(12, "12")
	value, ok := m.Load(12)
	fmt.Println(value, ok)
}
