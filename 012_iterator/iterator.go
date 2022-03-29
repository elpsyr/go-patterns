package iterator

import "fmt"

type Aggregate interface {
	Iterator() Iterator
}

type Iterator interface {
	First()
	IsDone() bool
	Next() interface{}
}

// 代表一个集合，里面的元素在start和end之间
type Numbers struct {
	start, end int
}

func NewNumbers(start, end int) *Numbers {
	return &Numbers{
		start: start,
		end:   end,
	}
}

// 迭代器
type NumbersIterator struct {
	// 迭代器拥有集合
	numbers *Numbers
	next    int
}

// First 初始化
func (i *NumbersIterator) First() {
	i.next = i.numbers.start
}

func (i *NumbersIterator) IsDone() bool {
	return i.next > i.numbers.end
}

func (i *NumbersIterator) Next() interface{} {
	if !i.IsDone() {
		next := i.next
		i.next++
		return next
	}
	return nil
}

func (n *Numbers) Iterator() Iterator {
	return &NumbersIterator{
		numbers: n,
		next:    n.start,
	}
}

func IteratorPrint(i Iterator) {
	for i.First(); !i.IsDone(); {
		c := i.Next()
		fmt.Printf("%#v\n", c)
	}

}
