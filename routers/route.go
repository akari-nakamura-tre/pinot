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
	salesRepo repositories.SalesRepository
}

type Services struct {
	idPOSSvc services.IdPOSService
	salesSvc services.SalesService
}

type Handlers struct {
	idPOSHdlr handlers.IdPOSHandler
	salesHdlr handlers.SalesHandler
}

func Init(e *echo.Echo, db *pinot.Connection) {
	repo := &Repositories{
		idPOSRepo: repositories.NewIdPOSRepository(db),
		salesRepo: repositories.NewSalesRepository(db),
	}
	svcs := &Services{
		idPOSSvc: services.NewIdPOSService(repo.idPOSRepo),
		salesSvc: services.NewSalesService(repo.salesRepo),
	}
	hdlrs := &Handlers{
		idPOSHdlr: handlers.NewIdPOSHandler(svcs.idPOSSvc),
		salesHdlr: handlers.NewSalesHandler(svcs.salesSvc),
	}
	e.GET("/idpos", hdlrs.idPOSHdlr.GetIdPOS)
	e.GET("/sales/summary", hdlrs.salesHdlr.GetSalesSummaryGroupByStoreAndDivision)
}
