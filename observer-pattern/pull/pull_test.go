package pull

import "testing"

func TestObserverPull(t *testing.T) {
	programBase := NewDevProgram()
	programBase.SetNewVersion("1.0")

	instanceProgma1 := NewInstanceProgram(programBase)
	instanceProgma2 := NewInstanceProgram(programBase)
	instanceProgma3 := NewInstanceProgram(programBase)

	programBase.AddUser(instanceProgma1)
	programBase.AddUser(instanceProgma2)
	programBase.AddUser(instanceProgma3)

	programBase.SetNewVersion("2.0")

	// запуск процедуры рассылки оповещения о возможности обновления
	programBase.NotifyForUpdate()

	// экземпляры программ сами применяют обновления при необходимости
	instanceProgma1.UpdateVersion()
	instanceProgma3.UpdateVersion()

	if instanceProgma1.version != instanceProgma3.version && instanceProgma3.version != "2.0" {
		t.Error("Expect pr1 and pr3 version", "2.0", ", but got", instanceProgma1.version, instanceProgma3.version)
	}

	if instanceProgma2.version != "1.0" {
		t.Error("Expect pr2 version", "1.0", ", but got", instanceProgma2)
	}
}