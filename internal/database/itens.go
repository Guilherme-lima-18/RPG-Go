// Package database é o catálogo estático de conteúdo do jogo: os itens e
// inimigos que EXISTEM no mundo, separados das regras que os manipulam.
//
// A separação importa: inventory define o que um equipamento É, database
// define QUAIS equipamentos existem. Balancear ou adicionar conteúdo mexe só
// aqui. A dependência é de mão única — database importa inventory, nunca o
// contrário.
package database

import "RPG/internal/inventory"

// Equipamentos indexa cada equipamento por ID, para consulta direta quando um
// baú, loja ou drop de inimigo só guardar o número do item.
//
// ATENÇÃO: a chave do mapa repete o campo ID de dentro do Item. Nada garante
// que continuem iguais ao editar. Uma função de validação na inicialização
// resolve, ou preencher o ID a partir da chave.
var Equipamentos = map[int]inventory.Equipamento{
	1: {
		Item: inventory.Item{
			ID:        1,
			Nome:      "Espada de ferro",
			Descricao: "Uma simples espada de ferro",
			Valor:     150,
			Peso:      5,
			Raridade:  inventory.Comum,
		},

		Ataque:       15,
		Durabilidade: 100,
		Slot:         inventory.Arma,
	},

	2: {
		Item: inventory.Item{
			ID:        2,
			Nome:      "Espada abençoada",
			Descricao: "Uma espada deixada para um viajante que almeja glória",
			Valor:     200,
			Peso:      20,
			Raridade:  inventory.Lendario,
		},

		Ataque:       45,
		Durabilidade: 1000,
		Slot:         inventory.Arma,
	},
}
