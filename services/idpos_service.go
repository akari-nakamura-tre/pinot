package services

import (
	"context"

	"github.com/akari-nakamura-tre/pinot/models"
	"github.com/akari-nakamura-tre/pinot/packages/util"
	"github.com/akari-nakamura-tre/pinot/repositories"
	"github.com/labstack/gommon/log"
	"github.com/startreedata/pinot-client-go/pinot"
)

type IdPOSService interface {
	GetIdPOS(ctx context.Context, limit int) []interface{}
}

type idPOSService struct {
	idPOSRepo repositories.IdPOSRepository
}

func NewIdPOSService(idPOSRepo repositories.IdPOSRepository) *idPOSService {
	return &idPOSService{idPOSRepo: idPOSRepo}
}

func (s *idPOSService) GetIdPOS(ctx context.Context, limit int) []interface{} {
	res, err := s.idPOSRepo.GetIdPOS(ctx, limit)
	if err != nil {
		log.Error(err)
	}
	jsonRes := util.CreateJSONResponse(res, mapIdPOSBrokerResponse)
	return jsonRes
}

func mapIdPOSBrokerResponse(table *pinot.ResultTable, row []interface{}) interface{} {
	return models.OutputData{
		Area:                   util.GetStringValue(table, row, "area"),
		AreaCode:               util.GetInt32Value(table, row, "areaCode"),
		BusinessDepartment:     util.GetStringValue(table, row, "businessDepartment"),
		BusinessDepartmentCode: util.GetInt32Value(table, row, "businessDepartmentCode"),
		CustomerCount:          util.GetInt32Value(table, row, "customerCount"),
		Date:                   util.GetStringValue(table, row, "date"),
		Department:             util.GetStringValue(table, row, "department"),
		DepartmentCode:         util.GetInt32Value(table, row, "departmentCode"),
		Division:               util.GetStringValue(table, row, "division"),
		DivisionCode:           util.GetInt32Value(table, row, "divisionCode"),
		Jan:                    util.GetStringValue(table, row, "jan"),
		Line:                   util.GetStringValue(table, row, "line"),
		LineCode:               util.GetInt32Value(table, row, "lineCode"),
		ProductName:            util.GetStringValue(table, row, "productName"),
		SalesQuantity:          util.GetInt32Value(table, row, "salesQuantity"),
		Store:                  util.GetStringValue(table, row, "store"),
		StoreCode:              util.GetInt32Value(table, row, "storeCode"),
		Team:                   util.GetStringValue(table, row, "team"),
		TeamCode:               util.GetInt32Value(table, row, "teamCode"),
		TotalPrice:             util.GetInt32Value(table, row, "totalPrice"),
		Zone:                   util.GetStringValue(table, row, "zone"),
		ZoneCode:               util.GetInt32Value(table, row, "zoneCode"),
	}
}
