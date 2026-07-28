// Package inventory define os itens e a mochila do jogador.
//
// Mesmo padrão do pacote enemy: Item é a BASE e cada categoria a embute,
// acrescentando só os campos próprios. São quatro categorias, uma por arquivo:
//
//	Equipamento -> equipamente.go  (arma/armadura, tem Slot e Durabilidade)
//	Consumivel  -> consumivel.go   (poção, gasta ao usar)
//	Material    -> material.go     (matéria-prima de crafting)
//	Quest       -> quest.go        (item de missão, não descartável)
//
// Raridade e Slot são tipos próprios (não int solto), definidos em
// raridade.go e slot.go — assim o compilador barra valores inválidos.
package inventory

// Item é o que TODO item tem, independente da categoria. Nunca é usado
// sozinho — existe para ser embutido.
//
// Peso é int, como todo valor numérico do jogo. Isso significa peso em
// unidades inteiras: a espada de ferro, que era 4.5, virou 5. Se um dia for
// preciso meio-quilo, o caminho é multiplicar a escala por 10 (guardar
// décimos) e não voltar para float.
type Item struct {
	ID        int
	Nome      string
	Descricao string
	Peso      int
	Valor     int
	Raridade  Raridade
}
