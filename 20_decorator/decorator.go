package decorator

type Component interface {
	Calc() int
}

type ConcreteComponent struct {
}

func (c *ConcreteComponent) Calc() int {
	return 0
}

type MulDecorator struct {
	Component
	num int
}

func WarpMulDecorator(c Component, num int) Component {
	return &MulDecorator{
		c,
		num,
	}
}
func (c *MulDecorator) Calc() int {
	return c.Component.Calc() * c.num
}

type AddDecorator struct {
	Component
	num int
}

func WarpAddDecorator(c Component, num int) Component {
	return &AddDecorator{
		c,
		num,
	}
}

func (d *AddDecorator) Calc() int {
	return d.Component.Calc() + d.num
}
