package esteservice
	
import (
	"context"

	"net/http"
)

type EsteRepository interface {
	GetAll(ctx context.Context) ([]any, error)
	Create(ctx context.Context, payload any) (any, error)
	Update(ctx context.Context, payload any, financialAccountId int) (any, error)
	FindByID(ctx context.Context, financialAccountId int) (any, error)
	Delete(ctx context.Context, financialAccountId int) error
	Active(ctx context.Context, financialAccountId int) error
}

type EsteService struct {
	repository EsteRepository
}

func NewEsteService(repository EsteRepository) *EsteService {
	return &EsteService{
		repository: repository,
	}
}

func (c *EsteService) GetAll(w http.ResponseWriter, r *http.Request) {
}

func (c *EsteService) Create(w http.ResponseWriter, r *http.Request) {
}

func (c *EsteService) FindByID(w http.ResponseWriter, r *http.Request) {
}

func (c *EsteService) Update(w http.ResponseWriter, r *http.Request) {
}

func (c *EsteService) Delete(w http.ResponseWriter, r *http.Request) {
}

func (c *EsteService) Active(w http.ResponseWriter, r *http.Request) {
}
