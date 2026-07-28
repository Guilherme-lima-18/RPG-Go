# RPG-Go

*A simple RPG I developed in Go for practice.*

Um RPG de terminal escrito em Go, feito para praticar a linguagem — com foco em
composição, interfaces e separação entre regras e conteúdo.

> **Status:** em desenvolvimento. O motor compila e é testado, mas `cmd/main.go`
> ainda está vazio: **não há como jogar ainda**. O que existe hoje são as regras
> de personagem, inimigos, itens e combate.

## Requisitos

- Go 1.26 ou superior

## Como rodar

```bash
git clone https://github.com/Guilherme-lima-18/RPG-Go.git
cd RPG-Go

go build ./...   # compila tudo
go test ./...    # roda os testes
go run ./cmd     # ponto de entrada (ainda não faz nada)
```

## Estrutura

```
cmd/main.go        ponto de entrada — menus e loop do jogo (a fazer)
internal/
  combat/          interface Combatente + ExecutarTurno
  player/          Personagem, Classe, fábrica NovoPersonagem
  enemy/           Enemy (base) + Dragao, Goblin, Esqueleto
  inventory/       Item (base) + Equipamento, Consumivel, Material, Quest
  database/        catálogo estático: quais itens e inimigos existem
  world/           mapa, baús e NPCs (a fazer)
```

Para um mapa mais detalhado do código, com o grafo de dependências e a lista de
pendências, veja [MAPA.md](MAPA.md).

## Ideias de projeto

### 1. Composição por struct embutida

Go não tem herança. No lugar dela, uma struct base é **embutida** nas
especializações, que herdam campos e métodos:

```go
type Dragao struct {
	Enemy            // herda Nome, Vida, Dano, Defesa e os métodos de combate
	Mana     int
	Elemento string
}
```

O mesmo padrão aparece em `inventory`, onde `Item` carrega o que todo item tem
(`Nome`, `Peso`, `Valor`, `Raridade`) e cada categoria embute e acrescenta o
que é seu.

### 2. Um contrato de combate para todos

`combat.Combatente` define quatro métodos — `Atacar`, `ReceberDano`,
`EstaVivo`, `ObterNome`. Tanto `*Personagem` quanto `*Enemy` os cumprem, então
os dois cabem na mesma variável e na mesma função de turno:

```go
func ExecutarTurno(atacante, defensor Combatente) {
	defensor.ReceberDano(atacante.Atacar())
}
```

Em Go a implementação de interface é implícita — basta ter os métodos. Por isso
o pacote `combat` **não importa** `player` nem `enemy`, ficando como folha da
árvore de dependências. Qualquer tipo novo que ganhe os quatro métodos (um NPC
hostil, um boss, um pet) entra em combate sem alterar uma linha de `combat`.

Cada lado declara uma garantia em tempo de compilação, que quebra o build na
hora se algum método sair de sintonia com a interface:

```go
var _ combat.Combatente = (*Personagem)(nil)
```

### 3. Conteúdo separado das regras

`inventory` define o que um equipamento **é**; `database` define **quais**
equipamentos existem. Balancear ou adicionar conteúdo mexe só no `database`.
A dependência é de mão única.

## Convenções

- **Todo valor numérico é `int`** — atributos de combate, dano, peso e cura.
  Foi o que permitiu um contrato de combate único. Se um dia precisar de
  fração, o caminho é mudar a escala (guardar décimos), não voltar para `float64`.
- **Fórmula de dano:** `dano - defesa`, com piso de 1, idêntica para os dois
  lados — assim defesa alta nunca torna ninguém imortal.
- Código, tipos e comentários em português.

## Próximos passos

- [ ] Loop principal e menus em `cmd/main.go`
- [ ] Métodos de `Inventario`: adicionar, remover, usar, equipar
- [ ] Atributos derivados (`AtaqueTotal`) em vez de somar bônus no atributo base
- [ ] Fábricas `CriarGoblin()` / `CriarEsqueleto()`
- [ ] Mapa, baús e NPCs em `world/`
- [ ] Level up e uso do ouro
- [ ] Save/load

## Licença

Sem licença definida.
