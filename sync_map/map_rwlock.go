package syncmap

import "sync"


type SyncMap[T comparable, U any] struct {
	data map[T]U
	rwLock sync.RWMutex
}

func NewSyncMap[T comparable, U any]() *SyncMap[T, U] {
	return &SyncMap[T, U]{
		data: make(map[T]U),
	}
}

// Set 写
func (sm *SyncMap[T, U]) Set(key T, val U) {
	sm.rwLock.Lock()
	defer sm.rwLock.Unlock()
	sm.data[key] = val
}

// 读
func (sm *SyncMap[T, U]) Get(key T) (U, bool) {
	sm.rwLock.RLock()
	defer sm.rwLock.RUnlock()

	val, ok := sm.data[key]
	return val, ok
}

// 删除
func (sm *SyncMap[T, U]) Delete(key T) {
	sm.rwLock.Lock()
	defer sm.rwLock.Unlock()

	delete(sm.data, key)
}

// 返回键值对数量
func (sm *SyncMap[T, U]) Len() int {
	sm.rwLock.RLock()
	defer sm.rwLock.RUnlock()

	return len(sm.data)
}