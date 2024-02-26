package main

import (
	"api-rest/routes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/swaggo/swag"
)

type swagger struct{}

// ReadDoc carrega o arquivo do swagger
func (s *swagger) ReadDoc() string {
	file, _ := os.Getwd()
	doc, _ := ioutil.ReadFile(fmt.Sprintf("%s/docs/%s", file, "swagger.yaml"))

	return string(doc)
}

func init() {
	swag.Register(swag.Name, &swagger{})
}

func main() {
	routes.HandleRequest(echo.New())
}
