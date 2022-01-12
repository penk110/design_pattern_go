package singleton

import (
	"sync"
	"testing"
)

const parallelCount = 10000

func TestSingleHungry(t *testing.T) {
	hungry1 := GetHungary()
	hungry2 := GetHungary()

	if hungry1 != hungry2 {
		t.Errorf("[TestSingleHungry] get hungry instance failed!\n")
		return
	}
	t.Logf("[TestSingleHungry] get hungry instance success!\n")
}

func TestParallelHungry(t *testing.T) {
	signal := make(chan struct{})
	wg := new(sync.WaitGroup)
	wg.Add(parallelCount)
	instances := [parallelCount]*Singleton{}
	for i := 0; i < parallelCount; i++ {
		go func(wg *sync.WaitGroup, index int) {
			<-signal
			instances[index] = GetHungary()
			wg.Done()
		}(wg, i)
	}

	// 关闭 channel 所有协程同时竞争
	close(signal)
	wg.Wait()

	for i := 1; i < parallelCount; i++ {
		// 循环判断是否存在不一致的instance
		if instances[i] != instances[i-1] {
			t.Errorf("[TestParallelHungry] instance is not equal\n")
		}
	}

	t.Logf("[TestParallelHungry] parallel get hungry instance success!\n")
}

func TestSingletonLazy(t *testing.T) {
	lazy1 := GetLazy()
	lazy2 := GetLazy()

	if lazy1 != lazy2 {
		t.Errorf("[TestSingletonLazy] get lazy instance failed!\n")
		return
	}
	t.Logf("[TestSingletonLazy] get lazy instance success!\n")
}

func TestParallelLazy(t *testing.T) {
	signal := make(chan struct{})
	wg := new(sync.WaitGroup)
	wg.Add(parallelCount)
	instances := [parallelCount]*Singleton{}
	for i := 0; i < parallelCount; i++ {
		go func(wg *sync.WaitGroup, index int) {
			<-signal
			instances[index] = GetLazy()
			wg.Done()
		}(wg, i)
	}

	// 关闭 channel 所有协程同时竞争
	close(signal)
	wg.Wait()

	for i := 1; i < parallelCount; i++ {
		// 循环判断是否存在不一致的instance
		if instances[i] != instances[i-1] {
			t.Errorf("[TestParallelLazy] instance is not equal\n")
		}
	}

	t.Logf("[TestParallelLazy] parallel get lazy instance success!\n")
}
