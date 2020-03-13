// Пример паттерна Observer (Наблюдатель), где данные поставляются наблюдателю методом проталкивания
package push

// Тип DevSoft предоставляет интерфейс с функциями, которые используются разработчиками для обновления своих программ
type DevSoft interface {
	AddUser(pr Program)
	SetNewVersion(version string)
	NotifyForUpdate()
}

// Тип Program предоставляет интерфейс с функциями, которые используются каждой конкретной программой для обновления
type Program interface {
	updateVersion(version string)
}

// Тип DevProgram реализует базу установленных экземпляров программы и содержит актуальные данные для программ
type DevProgram struct {
	progs []Program
	version string
}

// Тип InstanceProgram описывает структуру экземпляра программы
type InstanceProgram struct {
	version string // текущая версия, которую нужно поменять до новой
}

// Функция создает объект базы экземплярова программы
func NewDevProgram() *DevProgram {
	return &DevProgram{}
}

// Функция создает объект базы экземпляров программы
func NewInstanceProgram(dp *DevProgram) *InstanceProgram {
	if dp != nil {
		return &InstanceProgram{version: dp.version}
	}
	return nil
}

// Функция добавляет экземпляр программы к базе экземпляров
func (dp *DevProgram) AddUser(pr Program) {
	dp.progs = append(dp.progs, pr)
}

// Функция для установки новой версии программного продукта
func (dp *DevProgram) SetNewVersion(version string) {
	dp.version = version
}

// --- Реализация метода проталкивания ---
// Функция для каждого экземпляра программы делает обновление версии
func (dp *DevProgram) NotifyForUpdate() {
	for _, progma := range dp.progs {
		progma.updateVersion(dp.version)
	}
}

// Функция для обновления версии программы до более новой при получении уведомления о наличии обновления
func (ip *InstanceProgram) updateVersion(version string) {
	ip.version = version
}


