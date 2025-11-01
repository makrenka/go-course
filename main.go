package main

import (
	"fmt"
	"go-course/lessons"
)

func main() {
	fmt.Println(lessons.Attack())

	attackWithDamageBoost := lessons.DamageBoostDecorator(lessons.Attack)
	fmt.Println(attackWithDamageBoost())

	attackWithCriticalHit := lessons.CriticalHitDecorator(lessons.Attack)
	fmt.Println(attackWithCriticalHit())

	attackWithSlowEffect := lessons.SlowEffectDecorator(lessons.Attack)
	fmt.Println(attackWithSlowEffect())
}
