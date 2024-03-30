package entity

type HargaEmasAPI struct {
	Data []Data `json:"data"`
	Meta Meta   `json:"meta"`
}

type Meta struct {
	Url    string `json:"url"`
	Engine string `json:"engine"`
}

type Data struct {
	Sell float64 `json:"sell"`
	Buy  float64 `json:"buy"`
	Type string  `json:"type"`
}
