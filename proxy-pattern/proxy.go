package proxy_pattern

// Тип Subject предоставляет интерфейс для реального объекта и его суррогата
type Subject interface {
	DoSomething() string
}

// Proxy реализует суррогат - содержит указатель на оригинальный объект
type Proxy struct {
	toRealObject *RealObject
}

// RealObject реализует реальный объект
type RealObject struct {
}

// Функция реального объекта
func (ro *RealObject) DoSomething(str string) string  {
	return str
}

// Обращение к функции реального объекта через суррогат
func (p *Proxy) DoSomething(str string) string {
	if p.toRealObject == nil {
		// Если суррогат еще не ссылается на реальный объект, создать экземпляр реального объекта и сохранить ссылку
		p.toRealObject = &RealObject{}
	}
	// Обращение к методу реального объекта по ссылке
	return p.toRealObject.DoSomething(str)
}