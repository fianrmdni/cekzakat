package handlers

import (
	"cekzakat/entity"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func calculateZakat(ZakatType string, harta float64) float64 {
	nisab := getNisab(ZakatType)
	if harta >= nisab {
		zakat := harta * 0.025
		return zakat
	}
	return 0
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

	switch {
	case newZakat.ZakatType == "penghasilan":
		calculateZakat(newZakat.ZakatType, newZakat.JumlahZakat)
	case newZakat.ZakatType == "tabungan":
		calculateZakat(newZakat.ZakatType, newZakat.JumlahZakat)
	case newZakat.ZakatType == "emas":
		calculateZakat(newZakat.ZakatType, newZakat.JumlahZakat)
	case newZakat.ZakatType == "perak":
		calculateZakat(newZakat.ZakatType, newZakat.JumlahZakat)
	default:
		fmt.Println("default")
	}
	return c.JSON(http.StatusOK, newZakat)

}