package teste

import "fmt"

type Pessoa struct {
	nome string
	idade int
	email string
}

// Desse jeito, todos que estiverem no pacote main poderá usar a variável
// A variável pessoa2 está recebendo os tipos que a strutura Pessoa recebe
var pessoa2 Pessoa = Pessoa{}

func Falar(p Pessoa){
	fmt.Printf("A pessoa %v com o email %v está testandoa função falar!! \n", p.nome, p.email)
}

// Aqui você está atrelando a o comportamento dessa função na Pessoa
func (p Pessoa) Ouvir(){
	fmt.Printf("A pessoa %v com o email %v está testandoa função ouvir!! \n", p.nome, p.email)
}

func (p Pessoa) Dormir(){
	fmt.Printf("Função dormir!!")
}

func teste() {

	pessoa2.nome = "Bianca"
	pessoa2.email = "bibi.yumi@gmail.com"
	pessoa2.idade = 19

	pessoa1 := Pessoa{
		nome: "João",
		idade: 21,
		email: "joaooli17@hotmail.com",
	}

	fmt.Printf("Olá, eu sou %v, tenho %v, meu email é: %v \n", pessoa1.nome, pessoa1.idade, pessoa1.email)
	Falar(pessoa2)
	pessoa1.Ouvir()
	pessoa2.Dormir()

}