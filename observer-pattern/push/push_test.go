package push

import "testing"

func TestObserverPush(t *testing.T) {
	programBase := NewDevProgram()
	programBase.SetNewVersion("1.0")

	instanceProgma1 := NewInstanceProgram(programBase)
	instanceProgma2 := NewInstanceProgram(programBase)
	instanceProgma3 := NewInstanceProgram(programBase)

	programBase.AddUser(instanceProgma1)
	programBase.AddUser(instanceProgma2)
	programBase.AddUser(instanceProgma3)

	programBase.SetNewVersion("2.0")

	// запустить процедуру обновления экземпляров программ
	programBase.NotifyForUpdate()

	switch "1.0" {
		case instanceProgma1.version: t.Error("Expect program version", "2.0", ", but got", "1.0")
		case instanceProgma2.version: t.Error("Expect program version", "2.0", ", but got", "1.0")
		case instanceProgma3.version: t.Error("Expect program version", "2.0", ", but got", "1.0")
	}
}