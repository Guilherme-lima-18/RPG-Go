package inventory

// Raridade classifica o item. É um int com nome próprio (não um int solto),
// então o compilador impede passar um número qualquer onde se espera raridade.
//
// A ordem das constantes importa: por serem crescentes, dá para comparar
// direto — `if item.Raridade >= Epico`. Adicionar níveis novos só no FIM, ou
// renumerar com cuidado, para não bagunçar comparações existentes.
type Raridade int

// iota numera automaticamente a partir de 0: Comum=0, Incomum=1, e assim por
// diante. Adicionar uma linha no fim não exige tocar nas outras.
const (
	Comum Raridade = iota
	Incomum
	Raro
	Epico
	Lendario
)

// nomesRaridade é a tradução para exibição. Fica separado das constantes para
// que o texto mostrado ao jogador mude sem afetar a lógica.
var nomesRaridade = map[Raridade]string{
	Comum:    "Comum",
	Incomum:  "Incomum",
	Raro:     "Raro",
	Epico:    "Épico",
	Lendario: "Lendário",
}

// String faz Raridade se imprimir sozinha em fmt.Println e afins, em vez de
// aparecer como número. É a interface fmt.Stringer, da biblioteca padrão —
// mesmo princípio de combat.Combatente: basta ter o método.
func (r Raridade) String() string {
	if nome, existe := nomesRaridade[r]; existe {
		return nome
	}
	return "Desconhecida"
}
