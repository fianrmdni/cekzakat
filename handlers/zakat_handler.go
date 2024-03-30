package handlers

import (
	"cekzakat/entity"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

func calculateZakat(ZakatType string, harta float64) (float64, string, error) {
	hargaEmas, err := CallHargaEmasAPI()
	if err != nil {
		return 0, "", errors.New("gagal mendapatkan harga emas: " + err.Error())
	}
	nisab := getNisab(ZakatType, hargaEmas)
	if nisab < 0 || harta < 0 {
		return 0, "", errors.New("nisab tidak terpenuhi / harta tidak sesuai")
	}
	if harta >= nisab {
		zakat := harta * 0.025
		return zakat, "Perhitungan sukses! hartamu mencapai nisab, kamu wajib bayar zakat", nil
	}
	return 0, "", nil
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

	if !IsLogin(newZakat.Username) {
		errNotLogin := new(entity.ResponseZakat)
		errNotLogin.Message = "user belum login / tidak terdaftar"

		return c.JSON(http.StatusInternalServerError, errNotLogin)
	}

	response, err := calculateZakat(newZakat.ZakatType, newZakat.JumlahZakat)
	if err != nil {
		errZakat := entity.ResponseZakat{
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, errZakat)
	}
	responseZakat := entity.ResponseZakat{
		Zakat: response,
	}

	return c.JSON(http.StatusOK, responseZakat)

}
