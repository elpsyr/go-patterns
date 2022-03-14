package adapter

import "fmt"

/*
假设我现在有一个运维系统，需要分别调用阿里云和 AWS 的 SDK 创建主机，
两个 SDK 提供的创建主机的接口不一致，此时就可以通过适配器模式，将两个接口统一。
PS：AWS 和 阿里云的接口纯属虚构，没有直接用原始的 SDK，只是举个例子
*/

/*
aws    sdk 提供的接口为RunInstance
aliyun sdk 提供的接口为CreateServer
*/

type ICreateServer interface {
	CreateServer(cpu, mem float64) error
}

// AWSClient aws sdk
type AWSClient struct {
}

func (c *AWSClient) RunInstance(cpu, mem float64) error {
	fmt.Printf("aws client run success ,cpu: %f,mem: %f", cpu, mem)
	return nil
}

// AwsClientAdapter 适配器
type AwsClientAdapter struct {
	Client AWSClient
}

func (a *AwsClientAdapter) CreateServer(cpu, mem float64) error {
	a.Client.RunInstance(cpu, mem)
	return nil
}

//----------------------------------

type AliyunClient struct {
}

func (a *AliyunClient) CreateServer(cpu, mem int) error {
	fmt.Printf("aliyun client run success ,cpu: %d,mem: %d", cpu, mem)
	return nil
}

//
type AliyunClientAdapter struct {
	Client AliyunClient
}

func (a *AliyunClientAdapter) CreateServer(cpu, mem float64) error {
	a.Client.CreateServer(int(cpu), int(mem))
	return nil
}
