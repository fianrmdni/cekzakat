package handlers

import (
	"cekzakat/entity"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

var url = "https://logam-mulia-api.vercel.app/prices/hargaemas-com"

func CallHargaEmasAPI() (float64, error) {
	hargaEmas := entity.HargaEmasAPI{}
	response, err := http.Get(url)
	if err != nil {
		return 0, errors.New("error call API harga emas: " + err.Error())
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, errors.New("error baca body harga emas: " + err.Error())
	}

	err = json.Unmarshal(responseBody, &hargaEmas)
	if err != nil {
		return 0, errors.New("error unmarshal : " + err.Error())
	}

	if len(hargaEmas.Data) < 1 {
		return 0, errors.New("error harga emas tidak ditemukan")
	}

	return hargaEmas.Data[0].Buy, nil
}
