package transaction

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/rafaelcam/go-transaction-service/internal/transaction/domain"
	log "github.com/sirupsen/logrus"
	"net/http"
)

//Service transaction interface for DB access
//go:generate mockgen -source=routes.go -package=mock -destination=../../mock/gomock_service.go Service
type Service interface {
	GetTransactions() ([]domain.Transaction, error)
}

// Router structs represents Banks Handlers
type Router struct {
	e   *echo.Echo
	svc Service
}

// NewRouter is creating New Repository Transaction Router Handlers
func NewRouter(e *echo.Echo, db *sqlx.DB) *Router {
	return &Router{
		e:   e,
		svc: domain.NewService(domain.NewStore(db)),
	}
}

// Routes , all transaction routes
func (r *Router) Routes() {
	r.e.GET("/transactions", r.getTransactions)
}

func (r *Router) getTransactions(c echo.Context) error {
	transactions, err := r.svc.GetTransactions()
	if err != nil {
		log.Error(err)
		return err
	}

	return c.JSON(http.StatusOK, transactions)
}