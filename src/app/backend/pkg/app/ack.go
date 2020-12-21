package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

type Status struct {
	Message string `json:"message"`
}

func Response(resp *resty.Response) Status {
	if resp.StatusCode() != http.StatusOK {
		e := Status{}
		json.Unmarshal(resp.Body(), &e)
		return e
	}
	return Status{Message: ""}
}
func Error(err error) Status {
	return Status{Message: err.Error()}
}

type Gin struct {
	C *gin.Context
}

func (g *Gin) SendMessage(httpCode int, msg string) {
	g.C.JSON(httpCode, Status{Message: msg})
	return
}
func (g *Gin) Send(httpCode int, json interface{}) {
	g.C.JSON(httpCode, json)
	return
}

//url 검사
func (g *Gin) ValidateUrl(params []string) error {

	valid := validation.Validation{}

	for _, name := range params {
		valid.Required(g.C.Param(name), name)
	}

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			return errors.New(fmt.Sprintf("[%s]%s", err.Key, err.Error()))
		}
	}
	return nil

}