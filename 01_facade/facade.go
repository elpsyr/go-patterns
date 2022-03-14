package facade

import "fmt"

func NewAPI() API {
	return &apiImpl{}

}

type API interface {
	Test() string
}

type apiImpl struct {
	//a AModuleAPI
	//b BModuleAPI
	*aModuleImpl
	*bModuleImpl
}

func (a *apiImpl) Test() string {
	//aRet := a.a.TestA()
	//bRet := a.b.TestB()
	aRet := a.TestA()
	bRet := a.TestB()
	return fmt.Sprintf("%s\n%s", aRet, bRet)
}

//------------------------
//返回实现了A接口的实例
func NewAModuleAPI() AModuleAPI {
	return &aModuleImpl{}
}

type AModuleAPI interface {
	TestA() string
}
type aModuleImpl struct{}

func (a *aModuleImpl) TestA() string {
	return "A module running"
}

//------------------------
//返回实现了B接口的实例
func NewBModuleAPI() BModuleAPI {
	return &bModuleImpl{}
}

type BModuleAPI interface {
	TestB() string
}
type bModuleImpl struct{}

func (a *bModuleImpl) TestB() string {
	return "B module running"
}
