package routers

import (
	"github.com/akari-nakamura-tre/pinot/handlers"
	"github.com/akari-nakamura-tre/pinot/repositories"
	"github.com/akari-nakamura-tre/pinot/services"
	"github.com/labstack/echo/v4"
	"github.com/startreedata/pinot-client-go/pinot"
)

type Repositories struct {
	idPOSRepo repositories.IdPOSRepository
}

type Services struct {
	idPOSSvc services.IdPOSService
}

type Handlers struct {
	idPOSHdlr handlers.IdPOSHandler
}

func Init(e *echo.Echo, db *pinot.Connection) {
	repo := &Repositories{
		idPOSRepo: repositories.NewIdPOSRepository(db),
	}
	svcs := &Services{
		idPOSSvc: services.NewIdPOSService(repo.idPOSRepo),
	}
	hdlrs := &Handlers{
		idPOSHdlr: handlers.NewIdPOSHandler(svcs.idPOSSvc),
	}
	e.GET("/idpos", hdlrs.idPOSHdlr.GetIdPOS)
}
