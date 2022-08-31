# ProjectPatternsGo
Projeto para fins acadêmico da matéria Engenharia de Software III.

## Padrão Criacional - Factory Method
O Factory Method disponibiliza uma interface para criação de objetos que irá instanciar outros objetos (subclasses).
No diretório factory deste projeto podemos encontrar o código onde foi criado um objeto sem expor a sua lógica ao cliente. 
Essa interface única (Database interface) pode ser chamada por 2 bancos de dados diferentes (mongodb e sqlite) e o cliente não saberá.
(Ver o código para ter mais clareza)

## Padrão estrutural - Decorator
Nesse padrão fazemos um wrap de funcionalidades que já existem e injetamos no topo de outras.
O resultado será todas essas funcionalidades em uma só.
No diretório decorator adicionamos funcionalidades no topo de RomanceBook para que no final o valor total do livro seja a adição de tudo.
(Ver o código para ter mais clareza)

## Padrão comportamental - Iterator 
Como o próprio nome já nos fala, neste padrão é permitido a iteração sequencial através de uma estrutura sem expor seus detalhes lógicos internos.
(Ver o código para ter mais clareza)



#### Para rodar o código você precisa rodar no terminal `go run main.go` dentro do diretório principal desse projeto
