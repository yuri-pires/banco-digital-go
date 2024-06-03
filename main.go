package main

import (
	"fmt"

	"banco-digital/clientes"
	"banco-digital/contas"
)

func main() {
	clienteYuri := clientes.ClientePessoaFisica{"Yuri Pires", "04940108007", "Programador", "Rua 23"}

	contaDoYuri := contas.ContaCorrente{
		Titular:       clienteYuri,
		AgenciaConta:  5647,
		AgenciaDigito: 4,
	}

	status, mensagem := contaDoYuri.Depositar(45.55)
	fmt.Printf("%p \n", &contaDoYuri)
	fmt.Println("Depósito realizado:", status, mensagem)

	fmt.Println(contaDoYuri)

}
