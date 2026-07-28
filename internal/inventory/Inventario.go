package inventory

// Inventario é a mochila do jogador, guardada dentro de player.Personagem.
//
// A escolha aqui foi uma lista POR CATEGORIA em vez de uma lista única de
// interface. Vantagem: cada lista já vem com o tipo concreto, então dá para ler
// Equipamentos[0].Ataque direto, sem type assertion. Custo: percorrer "todos os
// itens" exige varrer as quatro listas.
//
// Quem cria: player.NovoPersonagem, com as quatro listas vazias.
//
// Ainda NÃO há métodos — hoje o Inventario é só um depósito passivo. Adicionar,
// Remover, Usar, Equipar e PesoAtual moram aqui quando forem escritos, e é isso
// que falta para o inventário virar algo jogável.
//
// Não há limite de peso: o campo PesoMax que a criação do personagem esperava
// nunca existiu nesta struct. Decidir se entra.
type Inventario struct {
	Equipamentos []Equipamento
	Consumiveis  []Consumivel
	Materiais    []Material
	Quests       []Quest
}
