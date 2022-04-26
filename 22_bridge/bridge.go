package bridge

import "fmt"

type Tools interface {
	Send(text, to string)
}

type HowToSend interface {
	SendMessage(text, to string)
}

type SMS struct {
}

func (S *SMS) Send(text, to string) {
	fmt.Printf("send %s  to %s via SMS\n", text, to)
}

func ViaSMS() Tools {
	return &SMS{}
}

type EMAIL struct {
}

func (S *EMAIL) Send(text, to string) {
	fmt.Printf("send %s  to %s via SMS\n", text, to)
}

func ViaEMAIL() Tools {
	return &EMAIL{}
}

type CommonSend struct {
	tool Tools
}

func (c *CommonSend) SendMessage(text, to string) {
	c.tool.Send(text, to)
}

func NewCommonSend(t Tools) *CommonSend {
	return &CommonSend{
		tool: t,
	}
}

type UrgencySend struct {
	tool Tools
}

func (c *UrgencySend) SendMessage(text, to string) {
	c.tool.Send(fmt.Sprintf("[Urgency]%s", text), to)
}

func NewUrgencySend(t Tools) *UrgencySend {
	return &UrgencySend{
		tool: t,
	}
}
