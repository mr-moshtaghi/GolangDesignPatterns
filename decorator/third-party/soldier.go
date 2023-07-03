package third_party

import (
	"log"
)

type InterfaceSoldier interface {
	Attack() int
	Defence() int
}

type BasicSoldier struct {
}

func (b BasicSoldier) Attack() int {
	//fmt.Println("attack in soldier with basic")
	return 1
}

func (b BasicSoldier) Defence() int {
	//fmt.Println("defence in soldier with shield")
	return 1
}

func DisplayStats(soldier InterfaceSoldier) {
	log.Printf("soldier stats: attack %d, defence %d", soldier.Attack(), soldier.Defence())
}
