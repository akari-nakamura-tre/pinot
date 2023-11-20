package handlers

import (
	"net/http"

	"github.com/akari-nakamura-tre/pinot/services"
	"github.com/labstack/echo/v4"
)

type SalesHandler interface {
	GetSalesSummaryGroupByStoreAndDivision(c echo.Context) error
}

type salesHandler struct {
	salesSvc services.SalesService
}

func NewSalesHandler(salesSvc services.SalesService) *salesHandler {
	return &salesHandler{
		salesSvc: salesSvc,
	}
}

func (h *salesHandler) GetSalesSummaryGroupByStoreAndDivision(c echo.Context) error {
	ctx := c.Request().Context()
	res := h.salesSvc.GetSalesSummaryGroupByStoreAndDivision(ctx)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": res,
	})
}
