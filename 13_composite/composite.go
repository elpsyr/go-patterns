package composite

import "fmt"

type Component interface {
	Parent() Component
	SetParent(Component)
	Name() string
	SetName(string)
	AddChild(Component)
	Print(string)
}

const (
	LeafNode = iota
	CompositeNode
)

func NewComponent(kind int, name string) Component {
	var c Component
	switch kind {
	case LeafNode:
		c = NewLeaf()
	case CompositeNode:
		c = NewComposite()
	}
	c.SetName(name)
	return c
}

type component struct {
	parent Component
	name   string
}

func (c *component) Parent() Component {
	return c.parent
}

func (c *component) SetParent(parent Component) {
	c.parent = parent
}

func (c *component) Name() string {
	return c.name
}

func (c *component) SetName(name string) {
	c.name = name
}

func (c *component) AddChild(child Component) {
}

func (c *component) Print(s string) {
}

type Leaf struct {
	component
}

func NewLeaf() *Leaf {
	return &Leaf{}
}

func (l *Leaf) Print(s string) {
	fmt.Printf("%s-%s\n", s, l.Name())
}

type Composite struct {
	component
	childes []Component
}

func NewComposite() *Composite {
	return &Composite{
		childes: make([]Component, 0),
	}
}
func (c *Composite) AddChild(child Component) {
	child.SetParent(c)
	c.childes = append(c.childes, child)
}

func (c *Composite) Print(pre string) {
	fmt.Printf("%s-%s\n", pre, c.Name())
	for _, comp := range c.childes {
		comp.Print(pre + pre)
	}
}
