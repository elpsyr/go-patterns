package command

import "fmt"

type Command interface {
	Execute()
}

// 可以理解为脚本库
type MotherBoard struct {
}

func (*MotherBoard) Start() {
	fmt.Println("system starting")
}

func (*MotherBoard) Reboot() {
	fmt.Println("system rebooting")
}

//开机脚本
type StartCommand struct {
	mb *MotherBoard
}

func NewStartCommand(mb *MotherBoard) *StartCommand {
	return &StartCommand{
		mb: mb,
	}
}

func (c *StartCommand) Execute() {
	c.mb.Start()
}

//重启脚本
type RebootCommand struct {
	mb *MotherBoard
}

func NewRebootCommand(mb *MotherBoard) *RebootCommand {
	return &RebootCommand{
		mb: mb,
	}
}

func (c *RebootCommand) Execute() {
	c.mb.Reboot()
}

type Box struct {
	button1 Command
	button2 Command
}

func NewBox(button1, button2 Command) *Box {
	return &Box{
		button1: button1,
		button2: button2,
	}

}

func (b *Box) PressButton1() {
	b.button1.Execute()

}

func (b *Box) PressButton2() {
	b.button2.Execute()
}
