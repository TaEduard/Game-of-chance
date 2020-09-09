package main

import (
	"github.com/TaEduard/GolangGame/utils"
	"fmt"
)

type Entity struct {
	name               string
	Health             float64
	Strength           int
	Defence            int
	Speed              int
	Luck               int
	DefenceSkillCd     int
	DefenceSkillActive bool
}

func (e *Entity) InitHero(name string) *Entity {
	e.name = name
	e.Health = float64(utils.RandomNoGen(70, 100))
	e.Strength = utils.RandomNoGen(70, 80)
	e.Defence = utils.RandomNoGen(45, 55)
	e.Speed = utils.RandomNoGen(40, 50)
	e.Luck = utils.RandomNoGen(10, 30)
	e.DefenceSkillCd = -2
	e.DefenceSkillActive = false
	fmt.Println("Hero Name:", e.name, " HP:", e.Health, " STR:", e.Strength, " DEF:", e.Defence, " Spd:", e.Speed, " LUK:", e.Luck)
	return e
}

func (e *Entity) InitVillan(name string) *Entity {
	e.name = name
	e.Health = float64(utils.RandomNoGen(60, 90))
	e.Strength = utils.RandomNoGen(60, 90)
	e.Defence = utils.RandomNoGen(40, 60)
	e.Speed = utils.RandomNoGen(40, 60)
	e.Luck = utils.RandomNoGen(25, 40)
	e.DefenceSkillCd = -9
	e.DefenceSkillActive = false
	fmt.Println("Villain Name:", e.name, " HP:", e.Health, " STR:", e.Strength, " DEF:", e.Defence, " Spd:", e.Speed, " LUK:", e.Luck)
	return e
}

type Battle struct {
	Hero    *Entity
	Villain *Entity
	Turn    int
	maxTurn int
}

func (b *Battle) Init(heroName string, villainName string, maxTurn int) {
	b.Turn = 0
	b.Hero = new(Entity).InitHero(heroName)
	b.Villain = new(Entity).InitVillan(villainName)
	b.maxTurn = maxTurn
	b.StartBattle()
}

func (b *Battle) StartBattle() {

	var first *Entity
	var second *Entity

	var winner *Entity

	if b.Hero.Speed > b.Villain.Speed {
		first = b.Hero
		second = b.Villain
	} else if b.Hero.Speed < b.Villain.Speed {
		first = b.Villain
		second = b.Hero
	} else if b.Hero.Speed == b.Villain.Speed {
		if b.Hero.Luck > b.Villain.Luck {
			first = b.Hero
			second = b.Villain
		} else if b.Hero.Luck < b.Villain.Luck {
			first = b.Villain
			second = b.Hero
		} else {
			fmt.Println("Not instructed on what to do in this case --- could do anything.")
		}
	}

	for i := 0; i < b.maxTurn; i++ {
		if first.Health <= 0 {
			winner = second
			break
		}
		if second.Health <= 0 {
			winner = first
			break
		}
		if i%2 == 0 {
			b.Attack(first, second)
		} else {
			b.Attack(second, first)
		}
	}

	if winner.DefenceSkillCd != -9 {
		fmt.Println("The winner is Hero ", winner.name)
	}else{
		fmt.Println("The winner is Villain ", winner.name)
	}

}

func (b *Battle) Attack(attacker *Entity, defender *Entity) {
	fmt.Println("\n Attacker:", attacker.name, " Defender: ", defender.name, " Turn", b.Turn)

	var locHero *Entity
	if attacker.DefenceSkillCd != -9 {
		locHero = attacker
	} else {
		locHero = defender
	}
	defence1 := utils.RandomNoGen(1, 5)
	defence2 := utils.RandomNoGen(1, 5)
	if defence1 == defence2 && locHero.DefenceSkillCd != b.Turn-1 {
		locHero.DefenceSkillActive = true
		locHero.DefenceSkillCd = b.Turn
		fmt.Println("SKILL: ", locHero.name, "Resilience Activated!")
	}

	damage := attacker.Strength - defender.Defence

	if attacker.DefenceSkillCd != -9 {
		chance1 := utils.RandomNoGen(1, 10)
		chance2 := utils.RandomNoGen(1, 10)
		if chance1 == chance2 {
			chance3 := utils.RandomNoGen(1, 100)
			chance4 := utils.RandomNoGen(1, 100)
			if chance3 == chance4 {
				fmt.Println("Attack: ", attacker.name, "landed a Super Critical Strike dealing ", damage*3, " damage.")
				defender.Health -= float64(damage * 3)
				fmt.Println(defender.name, " has ", defender.Health, "HP left.")
			}
			fmt.Println("Attack: ", attacker.name, "landed a Critical Strike dealing ", damage*2, " damage.")
			defender.Health -= float64(damage * 2)
			fmt.Println(defender.name, " has ", defender.Health, "HP left.")
		} else {
			fmt.Println("Attack: ", attacker.name, "landed a Normal Strike dealing ", damage, " damage.")
			defender.Health -= float64(damage)
			fmt.Println(defender.name, " has ", defender.Health, "HP left.")
		}
	} else {
		if defender.DefenceSkillActive == true {
			fmt.Println("Attack: ", attacker.name, "landed a Normal Strike dealing ", float64(damage)/float64(2), " damage.")

			defender.Health -= float64(damage) / float64(2)
		} else {
			fmt.Println("Attack: ", attacker.name, "landed a Normal Strike dealing ", damage, " damage.")
			defender.Health -= float64(damage)
		}
		fmt.Println(defender.name, " has ", defender.Health, "HP left.")
	}
	locHero.DefenceSkillActive = false
}

func main() {
	new(Battle).Init("Horia", "Boieru", 20)
}
