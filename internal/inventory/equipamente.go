package inventory

// Equipamento é o item vestível: soma Ataque/Defesa ao personagem enquanto
// estiver equipado. Slot define ONDE ele vai (arma, elmo, peito...), o que
// garante que só uma peça por lugar fique ativa.
//
// Durabilidade e NivelMinimo já estão modelados mas ainda não são lidos por
// ninguém — ganchos para desgaste em combate e restrição por nível.
//
// (o nome do arquivo tem typo: "equipamente" em vez de "equipamento";
// só afeta o disco, não o código.)
type Equipamento struct {
	Item
	Slot         Slot
	Ataque       int
	Defesa       int
	Durabilidade int
	NivelMinimo  int
}
