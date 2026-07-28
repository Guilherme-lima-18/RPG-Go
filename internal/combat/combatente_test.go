// Pacote de teste EXTERNO (combat_test, não combat): pode importar player e
// enemy sem risco de ciclo, já que os dois importam combat.
package combat_test

import (
	"testing"

	"RPG/internal/combat"
	"RPG/internal/enemy"
	"RPG/internal/player"
)

// TestJogadorEInimigoSaoIntercambiaveis é o teste que justifica a interface:
// os dois lados entram na MESMA variável combat.Combatente e passam pela mesma
// função de turno, nos dois sentidos.
func TestJogadorEInimigoSaoIntercambiaveis(t *testing.T) {
	heroi := player.NovoPersonagem("Teste", player.ObterClassesDisponiveis()[0])
	dragao := enemy.CriarDragaoVermelho()

	// A prova: uma fatia única guardando jogador e monstro lado a lado.
	lutadores := []combat.Combatente{heroi, dragao}

	for _, c := range lutadores {
		if !c.EstaVivo() {
			t.Errorf("%s deveria comecar vivo", c.ObterNome())
		}
		if c.Atacar() <= 0 {
			t.Errorf("%s deveria causar dano positivo", c.ObterNome())
		}
	}
}

// TestExecutarTurnoAplicaFormulaDeDano fixa a regra de dano: dano - defesa.
// O Paladino tem 15+5=20 de ataque; o dragão, 50 de defesa. 20-50 é negativo,
// então entra o piso de 1 — que é justamente o caso que impede defesa alta de
// tornar um combatente imortal.
func TestExecutarTurnoAplicaFormulaDeDano(t *testing.T) {
	heroi := player.NovoPersonagem("Teste", player.ObterClassesDisponiveis()[0])
	dragao := enemy.CriarDragaoVermelho()

	vidaAntes := dragao.Vida
	combat.ExecutarTurno(heroi, dragao)

	if perdido := vidaAntes - dragao.Vida; perdido != 1 {
		t.Errorf("piso de dano: esperava 1 de perda, veio %d", perdido)
	}

	// Sentido inverso, mesma função: dragão bate 80, defesa do Paladino é
	// 5+10=15, logo 65 de perda.
	vidaHeroi := heroi.Vida
	combat.ExecutarTurno(dragao, heroi)

	if perdido := vidaHeroi - heroi.Vida; perdido != 65 {
		t.Errorf("dano do dragao: esperava 65 de perda, veio %d", perdido)
	}
}

// TestEstaVivoAposDanoLetal garante a condição de parada do loop de batalha.
func TestEstaVivoAposDanoLetal(t *testing.T) {
	heroi := player.NovoPersonagem("Teste", player.ObterClassesDisponiveis()[0])

	for heroi.EstaVivo() {
		heroi.ReceberDano(1000)
	}

	if heroi.EstaVivo() {
		t.Error("heroi deveria estar morto")
	}
}
