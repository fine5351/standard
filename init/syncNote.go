package main

import (
	"fmt"
	"sync"
)

func SyncMapNote() {
	var syncMap sync.Map
	syncMap.Store("key1", "value1")
	syncMap.Store("key2", "value2")
	actual, loaded := syncMap.LoadOrStore("key3", "value3")
	if loaded {
		fmt.Printf("load key:%v, value:%v\n", "key3", actual)
	} else {
		fmt.Printf("store key:%v, value:%v\n", "key3", actual)
	}

	syncMap.Range(func(key, value interface{}) bool {
		fmt.Printf("range key:%v, value:%v\n", key, value.(string))
		return true
	})
}
