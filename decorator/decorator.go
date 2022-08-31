package decorator

//Nesse padrão fazemos um wrap de funcionalidades que já existem e injetamos no topo de outras.
//O resultado será todas as funcionalidades.
//No exemplo abaixo adicionamos os valores de 2 capitulos e 3 capítulos no topo do valor do preço do livro de Romance.
//O resultado final será o valor do livro de romance + 2 capitulos + 3 capitulos.
type Book interface {
	getPrice() int
}

type RomanceBook struct {
}

func (r *RomanceBook) getPrice() int {
	return 40
}

//Decorator 1
type TwoChapters struct {
	Book Book
}

func (t *TwoChapters) getPrice() int {
	bookPrice := t.Book.getPrice()
	return bookPrice + 10
}

//Decorator 2
type ThreeChapters struct {
	Book Book
}

func (t *ThreeChapters) GetPrice() int {
	bookPrice := t.Book.getPrice()
	return bookPrice + 20
}
