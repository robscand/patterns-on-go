package memento_pattern

import (
	"testing"
)

func TestMemento(t *testing.T) {
	data := new(Data)
	backupStorage := new(BackupStorage)

	data.Record = "This is my data"

	backupStorage.Storage = data.DoBackup()

	data.Record = "This is my new data"

	data.RestoreData(backupStorage.Storage)

	if data.Record != "This is my data" {
		// ошибка, если текущие данные не совпадают c изначальными
		t.Error("Expect data",  "\" This is my data \"", ", but got", "\"", data.Record, "\"")
	}
}
