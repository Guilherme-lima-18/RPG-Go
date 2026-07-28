// Package enemy define os inimigos.
//
// A ideia central: Enemy é a BASE com tudo que todo inimigo tem, e cada
// inimigo concreto (Dragao, Goblin, Esqueleto) EMBUTE Enemy e adiciona só o
// que é próprio dele. Ao embutir, o filho herda automaticamente os métodos de
// Enemy — ou seja, *Dragao já satisfaz combat.Combatente sem escrever nada.
//
// Para criar um inimigo novo: arquivo próprio, struct com `Enemy` embutido, e
// uma função CriarXxx() devolvendo o ponteiro já preenchido.
package enemy

import "RPG/internal/combat"

// Garantia em tempo de compilação de que *Enemy cumpre o contrato de combate.
// Não gera código nem custo em execução: se alguém mudar a assinatura de um
// método aqui ou na interface, o build quebra nesta linha, com mensagem clara,
// em vez de quebrar longe daqui na primeira batalha.
//
// Vale para todos os inimigos: como embutem Enemy, herdam os mesmos métodos.
var _ combat.Combatente = (*Enemy)(nil)

// Enemy é a base compartilhada, nunca usada sozinha no jogo — ela existe para
// ser embutida. Experiencia é o que o jogador ganha ao derrotá-lo.
type Enemy struct {
	Nome        string
	Vida        int
	Dano        int
	Defesa      int
	Experiencia int
}

// --- Implementação de combat.Combatente -----------------------------------
// Definidos UMA vez em Enemy e herdados por todos os inimigos que a embutem.
// Se um inimigo precisar de comportamento diferente (ex.: o Dragão usar sopro
// de fogo), basta declarar Atacar() no arquivo dele: o método do filho
// sobrescreve o da base.

func (e *Enemy) Atacar() int {
	return e.Dano
}

// ReceberDano usa a mesma fórmula do jogador (dano - defesa, mínimo 1),
// mantendo o combate simétrico entre os dois lados.
func (e *Enemy) ReceberDano(qnt int) {
	danoFinal := qnt - e.Defesa
	if danoFinal < 1 {
		danoFinal = 1
	}
	e.Vida -= danoFinal
}

func (e *Enemy) EstaVivo() bool {
	return e.Vida > 0
}

func (e *Enemy) ObterNome() string {
	return e.Nome
}
