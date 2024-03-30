package entity

type Zakat struct {
	Username    string
	ZakatType   string
	JumlahZakat float64
}

type ResponseZakat struct {
	Zakat   float64
	Message string
}
