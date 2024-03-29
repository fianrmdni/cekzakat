package handlers

import (
	"cekzakat/entity"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

func calculateZakat(ZakatType string, harta float64) (float64, error) {
	hargaEmas, err := CallHargaEmasAPI()
	if err != nil {
		return 0, errors.New("gagal mendapatakan harga emas " + err.Error())
	}
	nisab := getNisab(ZakatType, hargaEmas)
	if nisab < 0 || harta < 0 {
		return 0, errors.New("nisab tidak terpenuhi / harta tidak sesuai")
	}
	if harta >= nisab {
		zakat := harta * 0.025
		return zakat, nil
	}
	return 0, nil
}

func getNisab(ZakatType string, hargaEmas float64) float64 {
	if ZakatType == "penghasilan" {
		return 0
	} else if ZakatType == "tabungan" {
		return hargaEmas * 85
	} else if ZakatType == "emas" {
		return 85
	} else if ZakatType == "perak" {
		return 595
	} else {
		return -1
	}
}

func ZakatHandler(c echo.Context) error {
	newZakat := new(entity.Zakat)
	if err := c.Bind(newZakat); err != nil {
		return err
	}
	
	if !IsLogin(newZakat.Username){
		errNotLogin := entity.RensponseZakat{
			Message: "user belum login / tidak terdaftar",
		}
		return c.JSON(http.StatusInternalServerError, errNotLogin)
	}

	response, err := calculateZakat(newZakat.ZakatType, newZakat.JumlahZakat)
	if err != nil {
		errZakat := entity.RensponseZakat{
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, errZakat)
	}
	responseZakat := entity.RensponseZakat{
		Zakat: response,
	}

	return c.JSON(http.StatusOK, responseZakat)

}
