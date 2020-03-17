package adapter_pattern

import (
	"testing"
)

func TestAdapter(t *testing.T) {

	cat := new(Cat)
	dog := new(Dog)

	catSayResult := cat.Meow()
	dogSayResult := dog.Bark()

	dogHowCat := NewDog(cat)
	catHowDog := NewCat(dog)

	if catSayResult != dogHowCat.Meow() {
		t.Error("Expect dog and cat says \"meow\".")
	}
	if dogSayResult != catHowDog.Bark() {
		t.Error("Expect dog and cat says \"woof\".")
	}
}
