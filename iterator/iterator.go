package iterator

type Number struct {
	Value int
}
//Collection é apenas um container para objetos.
type Collection interface {
	createIterator() Iterator
}

type NumberCollection struct {
	Numbers []*Number
}

//Encapsula uma iteração, esse padrão depende dessa interface chamada Iterator.
//O método hasNext, observa se há elementos a serem iterados e retorna o próximo objeto na iteração.
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
