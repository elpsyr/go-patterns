package observer

import "testing"

func TestObserver(t *testing.T) {
	subject := NewSubject()
	r1 := NewReader("reader1")
	r2 := NewReader("reader2")
	r3 := NewReader("reader3")
	subject.Attach(r1)
	subject.Attach(r2)
	subject.Attach(r3)

	subject.UpdateContext("observer mode")

}
