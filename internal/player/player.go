// Package player define o personagem controlado pelo jogador.
//
// Peças deste arquivo, na ordem:
//  1. Classe          -> arquétipo escolhido na criação (só bônus, sem lógica)
//  2. AcoesPersonagem -> contrato de combate que Personagem cumpre
//  3. Personagem      -> o estado do jogador em si
//  4. NovoPersonagem  -> fábrica: base + bônus da classe
package player

import (
	"RPG/internal/inventory"
)

// Classe é só uma tabela de bônus. Ela NÃO tem comportamento próprio:
// os valores são somados uma única vez em NovoPersonagem e depois vivem
// dentro do Personagem. Criar uma classe nova = adicionar uma linha em
// ObterClassesDisponiveis, nada mais.
type Classe struct {
	Nome        string
	BonusVida   int
	BonusMana   int
	BonusAtaque int
	BonusDefesa int
}

// AcoesPersonagem é o contrato de combate do lado do jogador.
//
// Existe um contrato espelhado em enemy.AcoesEnemy. Hoje os dois NÃO são
// intercambiáveis porque este usa int e o do inimigo usa float64 — quando
// unificarmos tudo em int, os dois viram uma interface só (algo como
// combat.Combatente) e o loop de batalha passa a tratar jogador e inimigo
// pelo mesmo tipo.
type AcoesPersonagem interface {
	Atacar() int
	ReceberDano(quantidade int)
	EstaVivo() bool
	ObterNome() string
}

// Personagem guarda TODO o estado do jogador. É sempre usado como *Personagem:
// os métodos abaixo têm receptor de ponteiro porque alteram Vida, Ouro, etc.
type Personagem struct {
	Nome        string
	Vida        int
	Mana        int
	Experiencia int
	Level       int
	Ouro        int
	Defesa      int
	Ataque      int
	Inventario  inventory.Inventario
	Tipo        Classe
}

// --- Implementação de AcoesPersonagem -------------------------------------
// Os quatro métodos abaixo são o que faz *Personagem satisfazer a interface.
// Em Go isso é implícito: não existe "implements", basta ter os métodos.

// Atacar devolve o dano bruto. A subtração da defesa acontece do outro lado,
// dentro de ReceberDano do alvo.
func (p *Personagem) Atacar() int {
	return p.Ataque
}

// ReceberDano aplica a fórmula de dano: dano - defesa, com piso de 1.
// O piso existe para que defesa alta nunca deixe o personagem imortal.
func (p *Personagem) ReceberDano(qnt int) {
	danoFinal := qnt - p.Defesa
	if danoFinal < 1 {
		danoFinal = 1
	}
	p.Vida -= danoFinal
}

func (p *Personagem) EstaVivo() bool {
	return p.Vida > 0
}

func (p *Personagem) ObterNome() string {
	return p.Nome
}

// ObterClassesDisponiveis é o catálogo de classes, usado no menu de criação
// de personagem. Ponto único de edição para balancear ou adicionar classes.
func ObterClassesDisponiveis() []Classe {
	return []Classe{
		{Nome: "Paladino", BonusVida: 40, BonusMana: 20, BonusAtaque: 5, BonusDefesa: 10},
		{Nome: "Mago", BonusVida: 10, BonusMana: 80, BonusAtaque: 15, BonusDefesa: 5},
		{Nome: "Gerreiro", BonusVida: 60, BonusMana: 0, BonusAtaque: 10, BonusDefesa: 20},
	}
}

// NovoPersonagem é a fábrica do jogador: pega os valores base (iguais para
// todo mundo) e soma os bônus da classe escolhida. Depois daqui a Classe
// vira só um rótulo guardado em Tipo — ela não é mais consultada em combate.
//
// ATENÇÃO (pendência conhecida): o literal de inventory.Inventario abaixo usa
// os campos Itens e PesoMax, que NÃO existem na struct atual (ela tem
// Equipamentos, Consumiveis, Materiais e Quests). Este arquivo não compila
// enquanto isso não for alinhado — decidir se Inventario ganha peso máximo ou
// se a criação passa a inicializar as quatro listas.
func NovoPersonagem(nomeEscolhido string, classeEscolhida Classe) *Personagem {

	// Valores base, antes de qualquer bônus. Balanceamento global fica aqui.
	vidaBase := 100
	ataqueBase := 15
	defesaBase := 5
	manaBase := 50

	return &Personagem{
		Nome:        nomeEscolhido,
		Experiencia: 0,
		Level:       1,
		Ouro:        10,

		Vida:   vidaBase + classeEscolhida.BonusVida,
		Mana:   manaBase + classeEscolhida.BonusMana,
		Ataque: ataqueBase + classeEscolhida.BonusAtaque,
		Defesa: defesaBase + classeEscolhida.BonusDefesa,
		Tipo:   classeEscolhida,

		Inventario: inventory.Inventario{},
	}
}
