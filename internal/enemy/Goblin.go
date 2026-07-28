package enemy

// Goblin: inimigo comum, sem atributo extra nenhum — só a base.
// Falta a fábrica CriarGoblin() nos moldes de Dragao.go; enquanto ela não
// existir, não há como instanciar um goblin com valores definidos.
type Goblin struct {
	Enemy
}
