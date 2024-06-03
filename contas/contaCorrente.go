package contas

import "banco-digital/clientes"

type ContaCorrente struct {
	Titular       clientes.ClientePessoaFisica
	AgenciaConta  int
	AgenciaDigito int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorSaque float64) (bool, string) {
	validarValorDoSaque := valorSaque > 0 && valorSaque <= c.saldo

	if validarValorDoSaque {
		c.saldo -= valorSaque
		return true, "Saque efetuado"
	} else {
		return false, "Verifique a operação e tente novamente"
	}
}

func (c *ContaCorrente) Depositar(valorDeposito float64) (bool, string) {
	validarValorDeposito := valorDeposito > 0

	if validarValorDeposito {
		c.saldo += valorDeposito
		return true, "Depósito efetuado"
	}

	return false, "Informe um valor válido."
}
