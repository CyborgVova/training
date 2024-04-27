package main

import (
	"github.com/cyborgvova/training/oop/polimorfizm"
)

func main() {
	entity := polimorfizm.NewEntity(polimorfizm.NewCat("Barsik"))
	entity.MakeVoice()

	entity = polimorfizm.NewEntity(polimorfizm.NewDog("Sharik"))
	entity.MakeVoice()
}
