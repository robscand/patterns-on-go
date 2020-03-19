// Данная программа реализует проверку пациентов на симптомы коронавируса
package visitor_pattern

import "fmt"

// Тип CovidIdentifier предоставляет метод для обследования пациента на симптомы
type CovidIdentifier interface {
	Accept(doc CovidIdentifierDoctor)
}

//Тип CovidIdentifierVisitor предоставляет методы доктора (визитера) который знает технологию выявления симптомов
type CovidIdentifierDoctor interface {
	verifyFever(fev *Fever)
	verifyFatigue(fat *Fatigue)
	verifyDryСough(dc *DryСough)
	verifyAnorexia(an *Anorexia)
	verifyMyalgias(myal *Myalgias)
}

// Структура описывает характеристики симптомов Covid-19
// Лихорадка
type Fever struct {
	fever float32
}
// Усталость
type Fatigue struct {
	fatigue bool
}
// Сухой кашель
type DryСough struct {
	cough bool
}
// Потеря аппетита
type Anorexia struct {
	anorexia bool
}
// Боль в мышцах
type Myalgias struct {
	myalgias bool
}

// Реализация метода проверки всех симптомов доктором (визитером)
func (fev *Fever) Accept (doc CovidIdentifierDoctor) {
	doc.verifyFever(fev)
}
func (fat *Fatigue) Accept (doc CovidIdentifierDoctor) {
	doc.verifyFatigue(fat)
}
func (dc *DryСough) Accept (doc CovidIdentifierDoctor) {
	doc.verifyDryСough(dc)
}
func (an *Anorexia) Accept (doc CovidIdentifierDoctor) {
	doc.verifyAnorexia(an)
}
func (myal *Myalgias) Accept (doc CovidIdentifierDoctor) {
	doc.verifyMyalgias(myal)
}

// Структура Patient описывает человека, который может обладать симптомами Covid-19
type Patient struct {
	features []CovidIdentifier
}

// Функция GetCovid описывает развитие у пациента симптомов болезни (позволяет задать симптомы)
func GetCovid(fever float32, fatigue bool, cough bool, anorexia bool, myalgias bool) *Patient{
	p := new(Patient)
	p.features = []CovidIdentifier{
		&Fever{fever:fever},
		&Fatigue{fatigue:fatigue},
		&DryСough{cough:cough},
		&Anorexia{anorexia:anorexia},
		&Myalgias{myalgias:myalgias},
	}
	return p
}

// Реализация метода инициализирующего обследование пациента
func (p *Patient) Accept(doc CovidIdentifierDoctor) {
	for _, feature := range p.features {
		feature.Accept(doc)
	}
}

// Реализация конкретного доктора (визитера), проводящего обследование
type Doctor struct {
	FindFeatures uint8
	PatientViewResult []string
}

// Реализация методов по выявлению симптомов Covid-2019
func (doc *Doctor) verifyFever(fev *Fever) {
	if fev.fever > 36.6 {
		doc.PatientViewResult = append(doc.PatientViewResult, fmt.Sprintf("Patient has %f fever.\n", fev.fever))
		doc.FindFeatures++
	}
}
func (doc *Doctor) verifyFatigue(fat *Fatigue) {
	if fat.fatigue {
		doc.PatientViewResult = append(doc.PatientViewResult, fmt.Sprintf("Patient has fatigue.\n"))
		doc.FindFeatures++
	} else {
		doc.PatientViewResult = append(doc.PatientViewResult, fmt.Sprintf("Patient has no fatigue.\n"))
	}
}
func (doc *Doctor) verifyDryСough(dc *DryСough) {
	if dc.cough {
		doc.PatientViewResult = append(doc.PatientViewResult, fmt.Sprintf("Patient has dry cough.\n"))
		doc.FindFeatures++
	} else {
		doc.PatientViewResult = append(doc.PatientViewResult, fmt.Sprintf("Patient has no dry cough.\n"))
	}
}
func (doc *Doctor) verifyAnorexia(an *Anorexia) {
	if an.anorexia {
		doc.PatientViewResult = append(doc.PatientViewResult, fmt.Sprintf("Patient has anorexia.\n"))
		doc.FindFeatures++
	} else {
		doc.PatientViewResult = append(doc.PatientViewResult, fmt.Sprintf("Patient has no anorexia.\n"))
	}
}
func (doc *Doctor) verifyMyalgias(myal *Myalgias) {
	if myal.myalgias {
		doc.PatientViewResult = append(doc.PatientViewResult, fmt.Sprintf("Patient has myalgias.\n"))
		doc.FindFeatures++
	} else {
		doc.PatientViewResult = append(doc.PatientViewResult, fmt.Sprintf("Patient has no myalgias.\n"))
	}
}
