package inventory

// Material é matéria-prima: não se usa nem se equipa, serve de ingrediente
// para crafting e de moeda de troca com NPCs. Tipo agrupa por família
// ("Minério", "Erva") para as receitas filtrarem.
type Material struct {
	Item
	Tipo string
}
