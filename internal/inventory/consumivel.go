package inventory

// Consumivel é o item de uso único (poções). CuraHP e CuraMP são quanto ele
// devolve ao ser usado — zero num deles significa que aquele efeito não existe.
//
// Quatidade [sic: falta o "n"] é o empilhamento: em vez de guardar 5 poções
// iguais na lista, guarda-se uma entrada com Quatidade 5. Usar decrementa;
// chegando a zero, a entrada sai do inventário.
type Consumivel struct {
	Item
	CuraHP    int
	CuraMP    int
	Quatidade int
}
