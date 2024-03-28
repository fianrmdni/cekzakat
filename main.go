package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Zakat struct {
	zakatType   string
	jumlahZakat float64
}

func calculateZakat(zakatType string, harta float64) float64 {
	nisab := getNisab(zakatType)
	if harta >= nisab {
		zakat := harta * 0.025
		return zakat
	}
	return 0
}

func getNisab(zakatType string) float64 {
	if zakatType == "penghasilan" {
		return 0
	} else if zakatType == "tabungan" {
		return 1220000 * 85
	} else if zakatType == "emas" {
		return 85
	} else if zakatType == "perak" {
		return 595
	} else {
		return -1
	}
}

func perhitunganZakat(c echo.Context) error {
	newZakat := new(Zakat)
	if err := c.Bind(newZakat); err != nil {
		return err
	}

	switch {
	case newZakat.zakatType == "penghasilan":
		calculateZakat(newZakat.zakatType, newZakat.jumlahZakat)
	case newZakat.zakatType == "tabungan":
		calculateZakat(newZakat.zakatType, newZakat.jumlahZakat)
	case newZakat.zakatType == "emas":
		calculateZakat(newZakat.zakatType, newZakat.jumlahZakat)
	case newZakat.zakatType == "perak":
		calculateZakat(newZakat.zakatType, newZakat.jumlahZakat)
	default:
		fmt.Println("default")
	}
	return c.JSON(http.StatusOK, newZakat)

}

func main() {
	e := echo.New()

	e.POST("/perhitunganzakat", perhitunganZakat)
	e.Logger.Fatal(e.Start("localhost:1234"))

}
