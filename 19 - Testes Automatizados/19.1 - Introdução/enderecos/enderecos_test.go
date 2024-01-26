package enderecos_test

import (
	. "introducao-testes/enderecos"
	"testing"
)

// TESTE DE UNIDADE
// -v mostrará todos os processos feitos pelo teste
// --cover mostrará o cobrimento do seu teste sobre o arquivo
// --coverprofile (nome do arquivo desejado) mostrará o que não está sendo coberto em um arquivo
// go tool cover --func=nomeDoArquivoDesejado vai jogar no terminal a leitura do profile feito acima
// no lugar do func, pode ser usado --html para abrir um arquivo visual temporário no seu browser

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTiposDeEndereco(t *testing.T) {
	// TestNomeDaFunçãoTestada
	// O parallel vai fazer os testes rodarem em paralelo,
	// porém tem que ter nas funções de teste desejadas para rodar em paralelo
	t.Parallel()

	cenarioDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia do Imigrante", "Rodovia"},
		{"Praça das roças", "Tipo Inválido"},
		{"Estrada Qualquer", "Estrada"},
		{"RUA DOS BOBOS", "Rua"},
		{"AVENIDA REBOUÇAS", "Avenida"},
		{"", "Tipo Inválido"},
	}

	for _, cenario := range cenarioDeTeste {
		tipoDeEnderecoRecebido := TiposDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				tipoDeEnderecoRecebido,
				cenario.retornoEsperado,
			)
		}
	}
}

func TestQualquer(t *testing.T) {
	t.Parallel()

	if 1 > 2 {
		t.Error("Teste Quebrou!")
	}
}
