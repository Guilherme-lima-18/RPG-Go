package enemy

// Dragao é o exemplo mais completo do padrão de embedding: herda Nome, Vida,
// Dano, Defesa, Experiencia e os métodos de Enemy, e acrescenta o que só ele
// tem (Mana, sopro e elemento).
//
// SoproDeFogoDano ainda não é usado por nenhum método — é o gancho para um
// ataque especial. Quando existir, o jeito de ligá-lo é declarar Atacar()
// aqui: o método do Dragao passa a sobrescrever o de Enemy, e o combate
// continua chamando pela interface sem saber da diferença. (o nome presume
// fogo, mas há dragões de Gelo e Pedra — renomear junto com essa mudança.)
type Dragao struct {
	Enemy
	Mana            int
	SoproDeFogoDano int
	Elemento        string
}

// --- Fábricas -------------------------------------------------------------
// Cada CriarXxx devolve um dragão pronto. Elas são o "catálogo" de dragões:
// é aqui que se balanceiam os números, e não espalhado pelo código do jogo.

func CriarDragaoVermelho() *Dragao {
	return &Dragao{
		Enemy: Enemy{
			Nome:        "Dractus, o devorador",
			Vida:        1000,
			Dano:        80,
			Defesa:      50,
			Experiencia: 500,
		},
		Mana:            100,
		SoproDeFogoDano: 200,
		Elemento:        "Fogo",
	}
}

func CriarDragaoBranco() *Dragao {
	return &Dragao{
		Enemy: Enemy{
			Nome:        "Nerditor, o mogador",
			Vida:        1000,
			Dano:        80,
			Defesa:      50,
			Experiencia: 500,
		},
		Mana:            100,
		SoproDeFogoDano: 200,
		Elemento:        "Gelo",
	}
}

func CriarDragaoPreto() *Dragao {
	return &Dragao{
		Enemy: Enemy{
			Nome:        "Triviast, o destruidor",
			Vida:        1200,
			Dano:        100,
			Defesa:      50,
			Experiencia: 420,
		},
		Mana:            120,
		SoproDeFogoDano: 230,
		Elemento:        "Pedra",
	}
}
