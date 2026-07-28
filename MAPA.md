# Mapa do projeto

Orientação rápida para retomar o código sem reler tudo. Os detalhes estão nos
comentários de cada arquivo; aqui fica só o desenho geral.

## Estrutura

```
cmd/main.go            ponto de entrada — menus, loop do jogo (ainda vazio)
internal/
  combat/              interface Combatente + ExecutarTurno  <- contrato único
  player/player.go     Personagem, Classe, fábrica NovoPersonagem
  enemy/               Enemy (base) + Dragao, Goblin, Esqueleto
  inventory/           Item (base) + Equipamento, Consumivel, Material, Quest
                       e os tipos Raridade e Slot
  database/            catálogo estático: quais itens e inimigos existem
  world/               mapa, baús e NPCs (esqueleto, sem conteúdo)
```

Dependências (mão única, sem ciclos):

```
cmd  ->  player  ->  inventory
         player  ->  combat
         enemy   ->  combat
         database ->  inventory

combat não importa ninguém — é folha, de propósito.
```

## As duas ideias que se repetem

**1. Struct base embutida.** `Enemy` e `Item` não são usados sozinhos: existem
para serem embutidos. `Dragao` embute `Enemy` e já herda os métodos dele; se
precisar de comportamento próprio, basta redeclarar o método no filho, que
sobrescreve o da base. Mesma coisa em `Equipamento` sobre `Item`.

**2. Um contrato de combate só.** `combat.Combatente` tem quatro métodos —
`Atacar`, `ReceberDano`, `EstaVivo`, `ObterNome`. `*Personagem` e `*Enemy`
cumprem os dois, então cabem na mesma variável e na mesma função de turno.

Em Go não se declara "implements": ter os métodos basta. Por isso `combat` não
importa `player` nem `enemy` — e não deve. Cada lado tem uma linha de garantia
(`var _ combat.Combatente = (*Personagem)(nil)`) que quebra o build na hora se
algum método sair de sintonia.

Consequência prática: qualquer coisa nova que ganhe os quatro métodos — um NPC
hostil, um boss invocado, um pet — entra no combate sem alterar `combat`.

## Decisões tomadas

- **Todos os valores numéricos em `int`.** Vale para atributos de combate
  (`Vida`, `Mana`, `Ataque`, `Defesa`, `Experiencia`), para itens (`Ataque`,
  `Defesa`, `Durabilidade`, `CuraHP`, `CuraMP`) e também para `Peso`. Foi o que
  permitiu a interface única. Se um dia precisar de fração, a saída é mudar a
  escala (guardar décimos), não voltar para `float64`.
- **Fórmula de dano:** `dano - defesa`, com piso de 1, idêntica nos dois lados.

## Buracos conhecidos

- `cmd/main.go` vazio: sem menu, sem loop de jogo, sem criação de personagem.
- `world/` só tem os arquivos com o cabeçalho e a ideia anotada.
- Sem fábrica `CriarGoblin()` / `CriarEsqueleto()` — só o Dragão tem.
- `Inventario` não tem nenhum método (adicionar, remover, usar, peso total),
  e o limite de peso foi retirado da criação do personagem — decidir se volta.
- `Durabilidade` e `NivelMinimo` de `Equipamento` não são lidos por ninguém.
- `Dragao.SoproDeFogoDano` não é usado, e o nome presume fogo — mas existem
  dragões de Gelo e de Pedra.
- A chave do mapa `database.Equipamentos` duplica o campo `ID` do item, sem
  nada garantindo que continuem iguais.
- Nada de level up, ouro em uso, loja, save/load.

## Verificar

```
go build ./...
go test ./...     # internal/combat cobre a interface e a fórmula de dano
```
