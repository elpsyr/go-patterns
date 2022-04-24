package decorator

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	var c Component = &ConcreteComponent{}
	c = WarpAddDecorator(c, 10)
	c = WarpMulDecorator(c, 7)

	res := c.Calc()
	fmt.Println(res)
}
