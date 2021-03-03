package sword_and_shield

import "fmt"

type weapon interface {
	useWeapon(opponent *character)
}

type character struct {
	name   string
	hp     int
	level  int
	weapon weapon
}

func NewCharacter(name string, hp int) *character {
	return &character{
		name:  name,
		hp:    hp,
		level: 1,
	}
}

func (c *character) EquipWeapon(weapon weapon) {
	c.weapon = weapon
}

func (c *character) Attack(opponent *character) {
	fmt.Printf("%s ", c.name)
	c.weapon.useWeapon(opponent)
}
