package sword_and_shield

import "fmt"

type sword struct {
	damage int
	name   string
}

func NewSword(name string, damage int) *sword {
	return &sword{
		damage: damage,
		name:   name,
	}
}

func (s *sword) useWeapon(c *character) {
	fmt.Printf("đã chém %s bằng %s\n", c.name, s.name)
	c.hp = c.hp - s.damage
}

func (s *sword) getDamage() int {
	return s.damage
}
