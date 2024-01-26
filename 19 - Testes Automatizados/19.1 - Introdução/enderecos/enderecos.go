package enderecos

import "strings"

// TiposDeEndereco vai pegar o endereço e ver seu tipo (Rua, Avenida, Estrada) e o retorna
func TiposDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}
	enderecoEmLetraMinuscula := strings.ToLower(endereco)

	// string.split vai pegar a string e vai fazer um slice de acordo com o separador desejado
	// Nesse caso vai separar nos espaços
	// RUA DOS BOBOS
	// ["RUA", "DOS", "BOBOS"]
	// O index no final retornará a posição desejada desse split

	PrimeiraPalavraDoEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0]

	enderecoTemUmTipoValido := false

	for _, tipo := range tiposValidos {
		if tipo == PrimeiraPalavraDoEndereco {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		return strings.Title(PrimeiraPalavraDoEndereco)
	}

	return "Tipo Inválido"
}
