package task

import (
	"sync"
)

type store struct {
	sync.RWMutex
	lastinsertid 	uint64
	list			map[uint64]string
}

func NewStore() *store {
	return &store{
		lastinsertid: 	0,
		list:			make(map[uint64]string),
	}
}

func (s *store) Add(url string) (id uint64) {
	s.Lock()
	defer s.Unlock()

	s.lastinsertid++
	s.list[s.lastinsertid] = url
	return s.lastinsertid
}

func (s *store) List() []ResponseTask {
	s.RLock()
	defer s.RUnlock()

	list := make([]ResponseTask, len(s.list))
	listPos := 0
	for id, url := range s.list {
		list[listPos] = ResponseTask{
			Id: id,
			Url: url,
		}
		listPos++
	}

	return list
}

func (s *store) Delete(id uint64) bool {
	s.Lock()
	defer s.Unlock()

	if _, ok := s.list[id]; !ok {
		return false
	}

	delete(s.list, id)
	return true
}

func (s *store) Clear() {
	s.Lock()
	defer s.Unlock()
	s.list = make(map[uint64]string)
}

// Default store for tasks
var defaultStore = NewStore()

func StoreAdd(url string) uint64 {
	return defaultStore.Add(url)
}

func StoreList() []ResponseTask {
	return defaultStore.List()
}

func StoreDelete(id uint64) bool {
	return defaultStore.Delete(id)
}
