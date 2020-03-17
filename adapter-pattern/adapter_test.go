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

	if catSayResult != dogHowCat.CatMeow() {
		t.Error("Expect cat and dog says \"meow\".")
	}
	if dogSayResult != catHowDog.DogBark() {
		t.Error("Expect cat and dog says \"woof\".")
	}
}
