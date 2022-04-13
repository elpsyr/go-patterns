package state

import "testing"

func TestNewDayContext(t *testing.T) {
	context := NewDayContext()
	today := func() {
		context.Today()
		context.Next()
	}
	for i := 0; i < 7; i++ {
		today()
	}
}
