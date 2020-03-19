package visitor_pattern

import (
	"fmt"
	"testing"
)

func TestCovid(t *testing.T) {
	patient := GetCovid(37.7, true, true, true, true)
	doctor := new(Doctor)

	// пацинта начинает обследовать определенный доктор
	patient.Accept(doctor)
	fmt.Println(doctor.PatientViewResult)

	if doctor.FindFeatures != 5 {
		t.Errorf("Expect find 5 features of COVID-2019, but got %d\n.", doctor.FindFeatures)
	}
}