package estecontroller

import (
	"context"

	"net/http"
)

type EsteService interface {
	GetAll(context.Context) (any, error) // add your_response in any
	Create(context.Context, any) (any, error) // add your_response and request in any
	Update(context.Context, any, int) (any, error) // add your_response and request in any
	FindByID(context.Context, int) (any, error) // add your_response in any
	Delete(context.Context, int) error
	Active(context.Context, int) error
}

type EsteController struct {
	service EsteService
}

func NewEsteController(service EsteService) *EsteController {
	return &EsteController{
		service: service,
	}
}

func (c *EsteController) GetAll(w http.ResponseWriter, r *http.Request) {
}

func (c *EsteController) Create(w http.ResponseWriter, r *http.Request) {
}

func (c *EsteController) FindByID(w http.ResponseWriter, r *http.Request) {
}

func (c *EsteController) Update(w http.ResponseWriter, r *http.Request) {
}

func (c *EsteController) Delete(w http.ResponseWriter, r *http.Request) {
}

func (c *EsteController) Active(w http.ResponseWriter, r *http.Request) {
}
