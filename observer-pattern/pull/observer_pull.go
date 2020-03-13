// Пример паттерна Observer (Наблюдатель), где уведомления получаются методом вытягивания
package pull

// Тип DevSoft предоставляет интерфейс с функциями, которые используются разработчиками для управления своими программами
type DevSoft interface {
	AddUser(pr Program)
	SetNewVersion(version string)
	NotifyForUpdate()
	ReadVersion() string
}

// Тип Program предоставляет интерфейс с функциями, которые используются каждой конкретной программой для обновления
type Program interface {
	getNotify()
	UpdateVersion()
}

// Тип DevProgram реализует базу установленных экземпляров программы и содержит актуальные данные для программ
type DevProgram struct {
	progs []Program
	version string
}

// Тип InstanceProgram описывает структуру экземпляра программы вместе с адресом к общей базе
type InstanceProgram struct {
	toDev DevSoft // Адрес к базе актуальных данных разработчиков
	notify bool // уведомление о доступности новой версии
	version string // текущая версия, которую нужно поменять до новой
}

// Функция создает объект базы экземплярова программы
func NewDevProgram() *DevProgram {
	return &DevProgram{}
}

// Функция создает объект базы экземпляров программы
func NewInstanceProgram(dp *DevProgram) *InstanceProgram {
	if dp != nil {
		return &InstanceProgram{
			toDev : dp,
			version: dp.version,
		}
	}
	return nil
}

// Функция добавляет экземпляр программы к базе экземпляров
func (dp *DevProgram) AddUser(pr Program) {
	dp.progs = append(dp.progs, pr)
}

// Функция для установки новой версии
func (dp *DevProgram) SetNewVersion(version string) {
	dp.version = version
}

// --- Реализация метода вытягивания ---
// Функция для каждого экземпляра программы делает уведомление о том, что возможно обновление
func (dp *DevProgram) NotifyForUpdate() {
	for _, progma := range dp.progs {
		progma.getNotify() // рассылка уведомления о возможности обновления для каждого подключенного экземпляра
	}
}

// Функция для сохранения уведомления о возможности обновления
func (ip *InstanceProgram) getNotify() {
	ip.notify = true
}

// Функция для обновления версии программы до более новой при получении уведомления о наличии обновления
func (ip *InstanceProgram) UpdateVersion() {
	if ip.notify {
		// если есть обновление обновить и обнулить сообщение для последующих обновлений
		ip.version = ip.toDev.ReadVersion()
	}
}

// Функция для чтения данных из базы экземпляров программ
func (dp *DevProgram) ReadVersion() string {
	return dp.version
}


