package main

import (
	"fmt"

	"std/github.com/jpodlasnisky/aluragolang/OO_em_Go/contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {

	contaDenis := contas.ContaPoupanca{}
	contaPati := contas.ContaCorrente{}

	contaDenis.Depositar(100)
	fmt.Println(contaDenis.ObterSaldo())
	contaDenis.Sacar(10)
	fmt.Println(contaDenis.ObterSaldo())
	fmt.Println(contaDenis)
	fmt.Println(contaPati)

	contaPati.Depositar(500)
	fmt.Println("saldo Denis:", contaDenis.ObterSaldo())
	PagarBoleto(&contaDenis, 60)
	fmt.Println("Pagou boleto de 60, restou", contaDenis.ObterSaldo())
	fmt.Println("--------------------------------")
	fmt.Println("saldo Pati:", contaPati.ObterSaldo())
	PagarBoleto(&contaPati, 150)
	fmt.Println("Pagou boleto de 150, restou", contaPati.ObterSaldo())
	// contaExemplo := contas.ContaCorrente{}
	// contaExemplo.Depositar(100)

	// fmt.Println(contaExemplo.ObterSaldo())

	// clienteBruno := clientes.Titular{"Bruno", "123.456.789-09", "Desenvolvedor"}
	// contaBruno := contas.ContaCorrente{clienteBruno, 123, 123456, 100}

	// ==========================================================================
	// contaBruno := contas.ContaCorrente{Titular: clientes.Titular{
	// 	Nome:      "Bruno",
	// 	CPF:       "123.456.789-09",
	// 	Profissao: "Programador"},
	// 	NumeroAgencia: 0102, NumeroConta: 35050331, saldo: 100}

	// fmt.Println(contaBruno)
	// ==========================================================================
	//
	// contaDoJoao := contas.ContaCorrente{Titular: "João Pedro",
	// 	NumeroAgencia: 589,
	// 	NumeroConta:   123456,
	// 	saldo:         125.55}

	// fmt.Println("============================================================")
	// fmt.Println(contaDoJoao.saldo)
	// fmt.Println(contaDoJoao.Sacar(23.88))
	// fmt.Println("saldo atual da conta: R$", contaDoJoao.saldo)
	// status, valor := (contaDoJoao.Depositar(100.0))
	// fmt.Println(status)
	// fmt.Println("saldo atual da conta: R$", valor)
	// fmt.Println("============================================================")

	// conta01 := contas.ContaCorrente{Titular: "Julio", saldo: 300}
	// conta02 := contas.ContaCorrente{Titular: "Juliana", saldo: 100}

	// statusTransferencia := conta01.Transferir(200, &conta02)
	// fmt.Println(statusTransferencia)
	// fmt.Println(conta01)
	// fmt.Println(conta02)
}
