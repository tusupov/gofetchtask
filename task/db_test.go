package task

import (
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
)

var (
	testDB = NewStore()
	testUrl = "https://golang.org/"
)

func TestStore_Add(t *testing.T) {

	for i := 1; i <= 10; i++ {
		id := testDB.Add(testUrl)
		if id != uint64(i) {
			t.Fatalf("Expect %d, but %d", i, id)
		}
	}

	if testDB.lastinsertid != 10 {
		t.Fatalf("Expect %d, but %d", 10, testDB.lastinsertid)
	}

}

func TestStore_List(t *testing.T) {

	for _, responseTask := range testDB.List() {
		if responseTask.Url != testUrl {
			t.Fatalf("Expect [%s], but [%s]", testUrl, responseTask.Url)
		}
	}

}

func TestStore_Delete(t *testing.T) {

	for i := 1; i <= 5; i++ {
		ok := testDB.Delete(uint64(i))
		if !ok {
			t.Fatal("Should be deleted, but couldn't")
		}
	}

	len := len(testDB.List())
	if len != 5 {
		t.Fatalf("Length expect %d, but %d", 5, len)
	}

}

func BenchmarkStoreMutexMap_Add(b *testing.B) {

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			testDB.Add(testUrl)
		}
	})

}

func BenchmarkStoreSyncMap_Add(b *testing.B) {

	db := sync.Map{}
	id := uint64(0)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			newId := atomic.AddUint64(&id, 1)
			db.Store(newId, testUrl)
		}
	})

}

func BenchmarkStoreChannel_Add(b *testing.B) {

	db := make(map[uint64]string)
	id := uint64(0)
	keyCh := make(chan string, 1)
	go func() {
		for k := range keyCh {
			id++
			db[id] = k
		}
	}()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			keyCh <- testUrl
		}
	})

}

func BenchmarkStoreCAW_Add(b *testing.B) {

	db := make(map[uint64]string)
	id := uint64(0)
	lock := int32(0)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for {
				if lock == 0 && atomic.CompareAndSwapInt32(&lock, 0, 1) {
					break
				}
				runtime.Gosched()
			}
			id++
			db[id] = testUrl
			lock = 0
		}
	})

}

func BenchmarkStore_List(b *testing.B) {

	testDB.Clear()
	for i := 0; i < 10; i++ {
		testDB.Add(testUrl)
	}

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = testDB.List()
		}
	})

}

func BenchmarkStore_Delete(b *testing.B) {

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			testDB.Delete(1)
		}
	})

}
