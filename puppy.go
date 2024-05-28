package puppy

import (
	"github.com/ferdosian93/dog"
)

func Bark() string {
	return "hup!!"
}
func Barks() string {
	return "hup hupssss!!!"
}
func Mewo() string {
	return "mewooooooooooossssss"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}

func Hungry() string {
	return dog.WhenGetHungry(Mewo())
}
