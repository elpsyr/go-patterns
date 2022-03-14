package singleton

import "sync"

type Singleton interface {
	foo()
}

// 私有的，想要使用的对象
type singleton struct {
}

func (s singleton) foo() {
}

var (
	instance *singleton
	once     sync.Once
)

//返回一个私有类型的指针
func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

// 返回接口对调用者更友好
func GetInstanceV2() Singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
