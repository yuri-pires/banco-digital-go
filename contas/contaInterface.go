package contas

// ContaInterface define o contrato que uma conta deve seguir, especificando
// que ela deve implementar os métodos Sacar e Depositar.
type ContaInterface interface {
	Sacar(valorSaque float64) (bool, string)
	Depositar(valorDeposito float64) (bool, string)
}

func PagarBoleto(conta ContaInterface, valorBoleto float64) (bool, string) {
	status, mensagem := conta.Sacar(valorBoleto)
	return status, mensagem
}

// RealizarSaque tenta realizar um saque em uma conta que implementa a interface ContaInterface.
// A função verifica se a conta fornecida implementa a interface e, em caso afirmativo,
// chama o método Sacar da conta.
func RealizarSaque(conta ContaInterface, valorSaque float64) (bool, string) {
	status, mensagem := conta.Sacar(valorSaque)
	return status, mensagem
}

// RealizarDeposito tenta realizar um depósito em uma conta que implementa a interface ContaInterface.
// A função verifica se a conta fornecida implementa a interface e, em caso afirmativo,
// chama o método Depositar da conta.
func RealizarDeposito(conta ContaInterface, valorDeposito float64) (bool, string) {
	status, mensagem := conta.Depositar(valorDeposito)
	return status, mensagem
}
