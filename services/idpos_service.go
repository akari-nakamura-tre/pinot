package services

import (
	"context"
	"encoding/json"

	"github.com/akari-nakamura-tre/pinot/models"
	"github.com/akari-nakamura-tre/pinot/repositories"
	"github.com/labstack/gommon/log"
	"github.com/startreedata/pinot-client-go/pinot"
)

type IdPOSService interface {
	GetIdPOS(ctx context.Context, limit int) []models.OutputData
}

type idPOSService struct {
	idPOSRepo repositories.IdPOSRepository
}

func NewIdPOSService(idPOSRepo repositories.IdPOSRepository) *idPOSService {
	return &idPOSService{idPOSRepo: idPOSRepo}
}

func (s *idPOSService) GetIdPOS(ctx context.Context, limit int) []models.OutputData {
	res, err := s.idPOSRepo.GetIdPOS(ctx, limit)
	if err != nil {
		log.Error(err)
	}
	jsonRes := jsonResponse(res)
	return jsonRes
}

func getStringValue(table *pinot.ResultTable, row []interface{}, columnName string) string {
	for i, name := range table.DataSchema.ColumnNames {
		if name == columnName {
			return row[i].(string)
		}
	}
	return ""
}

func getInt32Value(table *pinot.ResultTable, row []interface{}, columnName string) int32 {
	for i, name := range table.DataSchema.ColumnNames {
		if name == columnName {
			val, _ := row[i].(json.Number)
			int64Val, _ := val.Int64()
			return int32(int64Val)
		}
	}
	return 0
}

func jsonResponse(brokerResp *pinot.BrokerResponse) []models.OutputData {
	var outputDataSlice []models.OutputData
	for _, row := range brokerResp.ResultTable.Rows {
		data := models.OutputData{
			Area:                   getStringValue(brokerResp.ResultTable, row, "area"),
			AreaCode:               getInt32Value(brokerResp.ResultTable, row, "areaCode"),
			BusinessDepartment:     getStringValue(brokerResp.ResultTable, row, "businessDepartment"),
			BusinessDepartmentCode: getInt32Value(brokerResp.ResultTable, row, "businessDepartmentCode"),
			CustomerCount:          getInt32Value(brokerResp.ResultTable, row, "customerCount"),
			Date:                   getStringValue(brokerResp.ResultTable, row, "date"),
			Department:             getStringValue(brokerResp.ResultTable, row, "department"),
			DepartmentCode:         getInt32Value(brokerResp.ResultTable, row, "departmentCode"),
			Division:               getStringValue(brokerResp.ResultTable, row, "division"),
			DivisionCode:           getInt32Value(brokerResp.ResultTable, row, "divisionCode"),
			Jan:                    getStringValue(brokerResp.ResultTable, row, "jan"),
			Line:                   getStringValue(brokerResp.ResultTable, row, "line"),
			LineCode:               getInt32Value(brokerResp.ResultTable, row, "lineCode"),
			ProductName:            getStringValue(brokerResp.ResultTable, row, "productName"),
			SalesQuantity:          getInt32Value(brokerResp.ResultTable, row, "salesQuantity"),
			Store:                  getStringValue(brokerResp.ResultTable, row, "store"),
			StoreCode:              getInt32Value(brokerResp.ResultTable, row, "storeCode"),
			Team:                   getStringValue(brokerResp.ResultTable, row, "team"),
			TeamCode:               getInt32Value(brokerResp.ResultTable, row, "teamCode"),
			TotalPrice:             getInt32Value(brokerResp.ResultTable, row, "totalPrice"),
			Zone:                   getStringValue(brokerResp.ResultTable, row, "zone"),
			ZoneCode:               getInt32Value(brokerResp.ResultTable, row, "zoneCode"),
		}
		outputDataSlice = append(outputDataSlice, data)
	}

	return outputDataSlice
}
