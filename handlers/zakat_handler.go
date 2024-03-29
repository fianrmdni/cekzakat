package handlers

import (
	"cekzakat/entity"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

func calculateZakat(ZakatType string, harta float64) (float64, error) {
	nisab := getNisab(ZakatType)

	if nisab < 0 || harta < 0 {
		return 0, errors.New("tipe zakat tidak sesuai / harta tidak sesuai")
	}
	if harta >= nisab {
		zakat := harta * 0.025
		return zakat, nil
	}
	return 0, nil
}

func getNisab(ZakatType string) float64 {
	if ZakatType == "penghasilan" {
		return 0
	} else if ZakatType == "tabungan" {
		return 1220000 * 85
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

	response, err := calculateZakat(newZakat.ZakatType, newZakat.JumlahZakat)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}


	responseZakat := entity.ZakatResponse{
		Zakat: response,
	}

	return c.JSON(http.StatusOK, responseZakat)

}