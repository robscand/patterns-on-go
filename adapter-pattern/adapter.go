// Данный пакет демонстрирует пример паттерна Адаптер
// Песели лают, а котики мяукают. Данный пример заставляет песеля мяукать, а котейку лаять.
package adapter_pattern

// Данный интерфейс предоставляет метод мяуканья
type HowCat interface {
	CatMeow() string
}

// Данный интерфейс предоставляет собачий лай
type HowDog interface {
	DogBark() string
}

// Класс реализует котика
type Cat struct {
	say string
}

// Класс реализует песеля
type Dog struct {
	say string
}

// Гав-метод
func (d *Dog) Bark() string {
	d.say = "woof"
	return d.say
}

// Мяу-метод
func (c *Cat) Meow() string {
	c.say = "meow"
	return c.say
}

// Класс-обертка содержит указатель на песеля
type DogEx struct {
	*Cat
}

// Класс-обертка содержит указатель на котейку
type CatEx struct {
	*Dog
}

// Конструкторы для оберток
func NewDog(c *Cat) HowCat {
	return &DogEx{c}
}

func NewCat(d *Dog) HowDog {
	return &CatEx{d}
}

// Методы для обращенных существ
func (td *DogEx) CatMeow() string {
	return td.Meow()
}

func (tc *CatEx) DogBark() string {
	return tc.Bark()
}




