package entity

type Zakat struct {
	ZakatType   string
	JumlahZakat float64
}

type ResponseZakat struct {
	Zakat   float64
	Message string
}
