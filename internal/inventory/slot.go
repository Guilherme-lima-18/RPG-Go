package inventory

// Slot é ONDE um equipamento é vestido. Serve de trava: só uma peça por slot
// fica ativa, então equipar um novo elmo desequipa o anterior automaticamente.
//
// Mesma técnica de Raridade — int com nome próprio, para o compilador barrar
// valores inválidos. Aqui a ORDEM não tem significado (não faz sentido dizer
// que Bota > Elmo), então slots novos podem entrar em qualquer posição.
type Slot int

const (
	Arma Slot = iota
	Escudo
	Elmo
	Peitoral
	Calca
	Bota
	Acessorio
)

// nomesSlot é a tradução para exibição na tela de equipamento.
var nomesSlot = map[Slot]string{
	Arma:      "Arma",
	Escudo:    "Escudo",
	Elmo:      "Elmo",
	Peitoral:  "Peitoral",
	Calca:     "Calça",
	Bota:      "Bota",
	Acessorio: "Acessório",
}

// String implementa fmt.Stringer, como em Raridade.
func (s Slot) String() string {
	if nome, existe := nomesSlot[s]; existe {
		return nome
	}
	return "Desconhecido"
}
