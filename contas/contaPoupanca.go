package contas

import "banco-digital-go/clientes"

type ContaPoupanca struct {
	Titular       clientes.ClientePessoaFisica
	AgenciaConta  int
	AgenciaDigito int
	Operacao      int
	saldo         float64
}

func (c *ContaPoupanca) Sacar(valorSaque float64) (bool, string) {
	validarValorDoSaque := valorSaque > 0 && valorSaque <= c.saldo

	if validarValorDoSaque {
		c.saldo -= valorSaque
		return true, "Saque efetuado"
	} else {
		return false, "Verifique a operação e tente novamente"
	}
}
