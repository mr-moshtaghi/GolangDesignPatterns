package main

import (
	"design-patterns/decorator/third-party"
)

func main() {
	basicSoldier := third_party.BasicSoldier{}
	third_party.DisplayStats(basicSoldier)

	soldierWithSword := SoldierWithSword{basicSoldier}
	third_party.DisplayStats(soldierWithSword)

	soldierWithShield := SoldierWithShield{basicSoldier}
	third_party.DisplayStats(soldierWithShield)

	soldierWithShieldWithSword := SoldierWithSword{
		soldier: SoldierWithShield{
			soldier: basicSoldier,
		},
	}
	third_party.DisplayStats(soldierWithShieldWithSword)
}

// Decorator 1: Soldier with  a sword

type SoldierWithSword struct {
	soldier third_party.InterfaceSoldier
}

func (s SoldierWithSword) Attack() int {
	//fmt.Println("attack in soldier with sword")
	return s.soldier.Attack() + 10
}

func (s SoldierWithSword) Defence() int {
	//fmt.Println("defence in soldier with sword")
	return s.soldier.Defence()
}

// Decorator 2: Soldier with Shield

type SoldierWithShield struct {
	soldier third_party.InterfaceSoldier
}

func (s SoldierWithShield) Attack() int {
	//fmt.Println("attack in soldier with shield")
	return s.soldier.Attack() - 6
}

func (s SoldierWithShield) Defence() int {
	//fmt.Println("defence in soldier with shield")
	return s.soldier.Defence() + 20
}
