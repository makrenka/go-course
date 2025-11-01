package lessons

import "math/rand/v2"

func Attack() string {
	return "Атака выполнена!"
}

func DamageBoostDecorator(attackFunc func() string) func() string {
	return func() string {
		return "Вам улыбнулась удача, нанесение урона увеличено на 10%! " + attackFunc()
	}
}

func CriticalHitDecorator(attackFunc func() string) func() string {
	return func() string {
		if rand.Float64() < 0.25 {
			return "Критический удар! Урон удвоен!" + attackFunc()
		}
		return attackFunc()
	}
}

func SlowEffectDecorator(attackFunc func() string) func() string {
	return func() string {
		return attackFunc() + "Цель замедлена на 2 хода!"
	}
}
