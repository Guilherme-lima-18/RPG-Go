package inventory

// Quest é item de missão. Missao guarda a qual missão ele pertence, e
// PodeDescartar é a trava de segurança: com false, nem venda nem descarte
// podem remover o item — evita o jogador travar a própria história.
type Quest struct {
	Item
	Missao        string
	PodeDescartar bool
}
