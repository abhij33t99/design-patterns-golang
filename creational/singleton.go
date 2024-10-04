package creational

import "sync"

type Singleton struct {
}

var instance *Singleton
var lock = &sync.Mutex{}

func getInstance() *Singleton {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()

		if instance == nil {
			instance = &Singleton{}
		}
	}

	return instance
}
