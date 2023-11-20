package services

import (
	"context"

	"github.com/akari-nakamura-tre/pinot/models"
	"github.com/akari-nakamura-tre/pinot/packages/util"
	"github.com/akari-nakamura-tre/pinot/repositories"
	"github.com/labstack/gommon/log"
	"github.com/startreedata/pinot-client-go/pinot"
)

type SalesService interface {
	GetSalesSummaryGroupByStoreAndDivision(ctx context.Context) []models.SalesSummary
}

type salesService struct {
	salesRepo repositories.SalesRepository
}

func NewSalesService(salesRepo repositories.SalesRepository) *salesService {
	return &salesService{
		salesRepo: salesRepo,
	}
}

func (s *salesService) GetSalesSummaryGroupByStoreAndDivision(ctx context.Context) []models.SalesSummary {
	res, err := s.salesRepo.GetSalesSummaryGroupByStoreAndDivision(ctx)
	jsonRes := createJsonResponse(res)
	if err != nil {
		log.Error(err)
	}
	return jsonRes
}

func createJsonResponse(brokerResp *pinot.BrokerResponse) []models.SalesSummary {
	var outputDataSlice []models.SalesSummary
	for _, row := range brokerResp.ResultTable.Rows {
		data := models.SalesSummary{
			StoreCode:    util.GetInt32Value(brokerResp.ResultTable, row, "storeCode"),
			Store:        util.GetStringValue(brokerResp.ResultTable, row, "store"),
			DivisionCode: util.GetInt32Value(brokerResp.ResultTable, row, "divisionCode"),
			Division:     util.GetStringValue(brokerResp.ResultTable, row, "division"),
			TotalPrice:   util.GetInt32Value(brokerResp.ResultTable, row, "totalPrice"),
		}
		outputDataSlice = append(outputDataSlice, data)
	}

	return outputDataSlice
}
