package memento_pattern

// Тип Data реализует данные (состояние), для которых требуется сохранение
type Data struct {
	Record string // какие-то данные, доступные для изменения
}

// Тип Backup реализует данные (состояние), которые сохранены и отделены от оригинального хранилища Data
type Backup struct {
	record string // те же данные, которые недоступны для непосредственного изменения
}

// Тип BackupStorage реализует хранилище на конкретный бэкап данных (указатель на Backup)
type BackupStorage struct {
	Storage *Backup
}

// Функция для создания бэкапа данных
func (d *Data) DoBackup() *Backup {
	// возвращение указателя на структуру Backup с дубликатом данных из Data
	return &Backup{record: d.Record}
}

// Функция для восстановления данных из бэкапа
func (d *Data) RestoreData(b *Backup) {
	// замещение данных в структуре Data данными из структуры Backup
	d.Record = b.GetBackupData()
}

// Функция для получения данных из бэкапа
func (b *Backup) GetBackupData() string {
	return b.record
}

