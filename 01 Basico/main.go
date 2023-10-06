package main

import (
	"net/http"

	"github.com/labstack/echo"
)

// import "src/src/"

// func main() {
// 	// test struct
// 	// var car objects.Car
// 	// car.Marca = "audi"
// 	// fmt.Println(car)

// 	// test interfaces
// 	// cuadrado := objects.Cuadrilatero{Base: 4, Altura: 4}
// 	// circulo := objects.Circulo{Radio: 5}

// 	// interfaces.CalcularArea(cuadrado)
// 	// interfaces.CalcularArea(circulo)

// 	//test channels

// 	src.Test()

// }

func main() {
	e := echo.New()
	//agregamos ruta
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
