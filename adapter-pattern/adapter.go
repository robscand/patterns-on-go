// Данный пакет демонстрирует пример паттерна Адаптер
// Песели лают, а котики мяукают. Данный пример заставляет песеля мяукать, а котейку лаять.
package adapter_pattern

// Данный интерфейс предоставляет метод мяуканья
type Meower interface {
	Meow() string
}

// Данный интерфейс предоставляет собачий лай
type Barker interface {
	Bark() string
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
type DogAdapter struct {
	*Cat
}

// Класс-обертка содержит указатель на котейку
type CatAdapter struct {
	*Dog
}

// Конструкторы для оберток
func NewDog(c *Cat) Meower {
	return &DogAdapter{c}
}

func NewCat(d *Dog) Barker {
	return &CatAdapter{d}
}

// Методы для обращенных существ
func (td *DogAdapter) Meow() string {
	return td.Cat.Meow()
}

func (tc *CatAdapter) Bark() string {
	return tc.Dog.Bark()
}




