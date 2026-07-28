// Package combat concentra as regras de batalha.
//
// A peça central é Combatente: o contrato ÚNICO que substituiu as duas
// interfaces espelhadas que existiam antes (player.AcoesPersonagem e
// enemy.AcoesEnemy). Com os dois lados falando em int, jogador e inimigo
// passaram a caber no mesmo tipo, e o combate não precisa mais saber quem é
// quem.
//
// Este pacote não importa player nem enemy — e não deve importar. Em Go a
// implementação de interface é implícita: basta ter os métodos. Manter combat
// como folha da árvore de dependências é o que permite a player e enemy
// importarem daqui sem criar ciclo.
package combat

// Combatente é tudo que pode lutar: entra em batalha, bate, apanha e morre.
//
// Qualquer tipo novo que cumpra estes quatro métodos (um NPC hostil, um boss
// invocado, um pet do jogador) entra no combate sem alterar nada aqui.
type Combatente interface {
	// Atacar devolve o dano BRUTO, antes da defesa do alvo.
	Atacar() int
	// ReceberDano aplica dano bruto: cada implementação desconta a própria
	// defesa e guarda o piso de 1.
	ReceberDano(quantidade int)
	// EstaVivo informa se ainda há vida. É a condição de parada do turno.
	EstaVivo() bool
	// ObterNome devolve o nome para exibição nas mensagens de batalha.
	ObterNome() string
}

// ExecutarTurno faz um combatente atacar o outro. É a função que justifica a
// interface: uma linha só, sem saber se atacante e defensor são jogador ou
// monstro — e serve para os dois sentidos do turno.
func ExecutarTurno(atacante, defensor Combatente) {
	defensor.ReceberDano(atacante.Atacar())
}
