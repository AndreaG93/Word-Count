package core

import (
	"container/list"
	"fmt"
	"sync"
)

var mutex = &sync.Mutex{}
var systemWorkersList = list.New()

func SubscribeSystemWorker(address string) {

	mutex.Lock()
	systemWorkersList.PushBack(address)

	for e := systemWorkersList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	mutex.Unlock()

}

// This function is used to retrieve System Worker Dictionary Network Address from a specified Worker ID.
func GetSystemWorkersDictionary() *list.List {
	return systemWorkersList
}

func GetAvailableWorker() int {

	return systemWorkersList.Len()

}
