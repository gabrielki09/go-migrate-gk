package esteroutes

import(
	"net/http"

	estecontroller "github.com/gabrielki09/go-scaffold-gk/internal/modules/financial_transfer/controller"
	esterepository "github.com/gabrielki09/go-scaffold-gk/internal/modules/financial_transfer/repository"
	esteservice "github.com/gabrielki09/go-scaffold-gk/internal/modules/financial_transfer/service"
	
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterEsteRoutes(r *http.ServeMux, db *pgxpool.Pool) {
	repo := esterepository.NewEsteRepository(db)
	service := esteservice.NewEsteService(repo)
	controller := estecontroller.NewEsteController(service)

	r.HandleFunc("GET /este", controller.GetAll)
	r.HandleFunc("GET /este/{id}", controller.FindByID)
	r.HandleFunc("POST /este", controller.Create)
	r.HandleFunc("PUT /este/{id}", controller.Update)
	r.HandleFunc("DELETE /este/delete/{id}", controller.Delete)
	r.HandleFunc("PATCH /este/active/{id}", controller.Active)
}
	