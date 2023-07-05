package human

import "log"

type Human struct {
	lifePoints int
}

func NewHuman() *Human {
	return &Human{lifePoints: 100}
}

func (h *Human) Display() {
	log.Print("life points: ", h.lifePoints)
}

func (h *Human) Damage(damagePoints int) {
	h.lifePoints = h.lifePoints - damagePoints
}
