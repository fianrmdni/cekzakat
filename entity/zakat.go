package entity

type Zakat struct {
	ZakatType   string
	JumlahZakat float64
}

type RensponseZakat struct {
	Zakat   float64
	Message string
}
