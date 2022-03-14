package adapter

type Target interface {
	Request() string
}

type Adaptee interface {
	SpecificRequest() string
}

// 目标是把 Adaptee的实现类变成Target的实现类
// 即调用Request方法，返回SpecificRequest的结果

type adapteeImpl struct {
}

func (a *adapteeImpl) SpecificRequest() string {
	return "原来的输出"
}

func NewAdaptee() Adaptee {
	return &adapteeImpl{}
}

//------------以上是旧的接口以及实现

//---------开始适配
type adapter struct {
	Adaptee
}

func (a *adapter) Request() string {
	return a.SpecificRequest()
}

func NewAdapter(adaptee Adaptee) Target {
	return &adapter{
		adaptee,
	}
}

// 总的来说，就是
// 包含一个旧接口，实现一个新接口
// 在新的接口实现里调用旧接口
