package main

import (
	"fmt"
	"reflect"
	"robertamadge/ProjectPatternsGo/decorator"
	"robertamadge/ProjectPatternsGo/factory"
	"robertamadge/ProjectPatternsGo/iterator"
)

//Main defines the two environments choosing between the two databases
func main() {
	//Factory
	env1 := "production"
	env2 := "development"

	//Database: mongo
	db1 := factory.DatabaseFactory(env1)
	//Database: sqlite
	db2 := factory.DatabaseFactory(env2)

	db1.PutData("Putdata", "this is mongodb")
	fmt.Println(db1.GetData("Getdata"))

	db2.PutData("Putdata", "this is sqlite")
	fmt.Println(db2.GetData("Getdata"))

	//Prints o tipo
	fmt.Println(reflect.TypeOf(db1).Name())
	//Prints a interface
	fmt.Println(reflect.TypeOf(&db1).Elem())

	fmt.Println(reflect.TypeOf(db2).Name())
	fmt.Println(reflect.TypeOf(&db2).Elem())



	//Decorator
	book := &decorator.RomanceBook{}

	//Adiciona 2 capítulos
	addTwoChapters := &decorator.TwoChapters{
		Book: book,
	}

	//Adiciona tres capítulos
	addThreeChapters := &decorator.ThreeChapters{
		Book: addTwoChapters,
	}
	//Prints o valor do livro com o preço adicional de 3 capítulos
	fmt.Println("The price of the book is: %d\n", addThreeChapters.GetPrice())



	//Iterator
	firstNumber := &iterator.Number{
		Value: 1,
	}

	secondNumber := &iterator.Number{
		Value: 2,
	}

	numberCollection := &iterator.NumberCollection{
		Numbers: []*iterator.Number{firstNumber, secondNumber},
	}

	itr := numberCollection.CreateIterator()
	for itr.HasNext() {
		number := itr.GetNext()
		fmt.Printf("Number is %+v\n", number)
	}
}
