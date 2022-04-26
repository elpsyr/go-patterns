package bridge

import "testing"

func TestNewCommonSend(t *testing.T) {
	s1 := NewCommonSend(ViaSMS())
	s1.SendMessage("have a drink", "bob")

	s2 := NewCommonSend(ViaEMAIL())
	s2.SendMessage("have a drink", "bob")

	s3 := NewUrgencySend(ViaSMS())
	s3.SendMessage("have a drink", "bob")

	s4 := NewUrgencySend(ViaEMAIL())
	s4.SendMessage("have a drink", "bob")
}
