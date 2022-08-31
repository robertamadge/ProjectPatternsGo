package iterator

type Number struct {
	Value int
}
type Collection interface {
	createIterator() Iterator
}

type NumberCollection struct {
	Numbers []*Number
}

type Iterator interface {
	HasNext() bool
	GetNext() *Number
}

type NumberIterator struct {
	index   int
	numbers []*Number
}

func (n *NumberCollection) CreateIterator() Iterator {
	return &NumberIterator{
		numbers: n.Numbers,
	}
}

func (n *NumberIterator) HasNext() bool {
	if n.index < len(n.numbers) {
		return true
	}
	return false
}

func (n *NumberIterator) GetNext() *Number {
	if n.HasNext() {
		number := n.numbers[n.index]
		n.index++
		return number
	}
	return nil
}
