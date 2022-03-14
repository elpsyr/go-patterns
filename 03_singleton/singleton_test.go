package singleton

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

const parCount = 100

func TestSingleTon(t *testing.T) {
	ins1 := GetInstance()
	//ins2 := GetInstance()
	ins3 := GetInstanceV2()

	fmt.Println(reflect.TypeOf(ins1))
	fmt.Println(reflect.TypeOf(ins3))
	if ins1 != ins3 {
		t.Fatalf("instance is not equal")

	}
}

func TestParallelSingleton(t *testing.T) {
	start := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(parCount)
	instances := [parCount]Singleton{}
	for i := 0; i < parCount; i++ {
		go func(index int) {
			<-start
			instances[index] = GetInstanceV2()
			wg.Done()
		}(i)
	}
	// 关闭channel，协程并行
	close(start)
	wg.Wait()
	for i := 1; i < parCount; i++ {
		if instances[i] != instances[i-1] {
			t.Fatalf("instance os not equal")

		}
	}

}
