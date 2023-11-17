package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/akari-nakamura-tre/pinot/services"
	"github.com/labstack/echo/v4"
)

type IdPOSHandler interface {
	GetIdPOS(c echo.Context) error
}

type idPOSHandler struct {
	idPOSSvc services.IdPOSService
}

func NewIdPOSHandler(idPOSSvc services.IdPOSService) *idPOSHandler {
	return &idPOSHandler{idPOSSvc: idPOSSvc}
}

func (h *idPOSHandler) GetIdPOS(c echo.Context) error {
	ctx := c.Request().Context()
	var l int
	limit := c.QueryParam("limit")
	if limit != "" {
		i, err := strconv.Atoi(limit)
		if err != nil {
			return fmt.Errorf("errror: %v", err)
		}
		l = i
	} else {
		l = 10
	}

	res := h.idPOSSvc.GetIdPOS(ctx, l)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": res,
	})
}
