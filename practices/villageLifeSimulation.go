package practices

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

type VillageElement interface {
	Update()
	FlushInfo() string
}

type Resident struct {
	Name    string
	Age     int
	Married bool
	Alive   bool
	Events  []string
}

type Animal struct {
	Name   string
	Age    int
	Type   string
	Alive  bool
	Events []string
}

type Village struct {
	Elements []VillageElement
}

func (r *Resident) AddYear() {
	r.Age++
}

func (a *Animal) AddYear() {
	a.Age++
}

func (r *Resident) Dead() {
	r.Alive = false
}

func (a *Animal) Dead() {
	a.Alive = false
}

func (r *Resident) ChangeMarried() {
	r.Married = !r.Married
}

func (r *Resident) Update() {
	r.AddYear()
	if rand.IntN(100) < 5 {
		r.Dead()
		r.Events = append(r.Events, "is dead")
	}
	if rand.IntN(100)+5 < 20 {
		r.ChangeMarried()
		if r.Married {
			r.Events = append(r.Events, "divorced")
		} else {
			r.Events = append(r.Events, "got married")
		}
	}
	if rand.IntN(100)+5 < 25 {
		r.Events = append(r.Events, "got a job")
	}
	if rand.IntN(100)+25 < 50 {
		r.Events = append(r.Events, "bought a car")
	}
}

func (a *Animal) Update() {
	a.AddYear()
	if rand.IntN(100) < 5 {
		a.Dead()
		a.Events = append(a.Events, "is dead")
	}
	if rand.IntN(100)+5 < 25 {
		a.Events = append(a.Events, "broke his paw")
	}
	if rand.IntN(100)+25 < 50 {
		a.Events = append(a.Events, "left home")
	}
}

func (v *Village) AddElement(e VillageElement) {
	v.Elements = append(v.Elements, e)
}

func (r *Resident) FlushInfo() string {
	info := fmt.Sprintf("Inhabitant %s (%d years old) is dead", r.Name, r.Age)
	if r.Alive {
		married := "not married"
		if r.Married {
			married = "married"
		}
		events := "none"
		if len(r.Events) > 0 {
			events = fmt.Sprint(strings.Join(r.Events, "\n"))
		}
		info = fmt.Sprintf("Inhabitant %s (%d years old), %s.\nEvents: %s\n", r.Name, r.Age, married, events)
	}

	r.Events = []string{}
	return info
}

func (a *Animal) FlushInfo() string {
	info := fmt.Sprintf("Animal %s (%s %d years old).", a.Name, a.Type, a.Age)
	a.Events = []string{}
	return info
}

func (v *Village) UpdateAll() {
	for _, v := range v.Elements {
		v.Update()
	}
}

func (v *Village) ShowAllInfo() string {
	info := ""
	for _, v := range v.Elements {
		info += v.FlushInfo()
	}
	return info
}
