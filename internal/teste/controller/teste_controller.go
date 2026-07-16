package testecontroller

import (
	"net/http"
)

type TesteService interface {
}

type TesteController struct {
	service TesteService
}

func NewTesteController(service TesteService) *TesteController {
	return &TesteController{
		service: service,
	}
}

func (c *TesteController) GetAll(w http.ResponseWriter, r *http.Request) {
}

func (c *TesteController) Create(w http.ResponseWriter, r *http.Request) {
}

func (c *TesteController) FindByID(w http.ResponseWriter, r *http.Request) {
}

func (c *TesteController) Update(w http.ResponseWriter, r *http.Request) {
}

func (c *TesteController) Delete(w http.ResponseWriter, r *http.Request) {
}
