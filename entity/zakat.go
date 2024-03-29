package entity

type Zakat struct {
	ZakatType   string
	JumlahZakat float64
}

type ZakatResponse struct {
	Zakat float64
	Message string
}