package repositories

import (
	"context"
	"fmt"

	"github.com/labstack/gommon/log"
	"github.com/startreedata/pinot-client-go/pinot"
)

type IdPOSRepository interface {
	GetIdPOS(ctx context.Context, limit int) (*pinot.BrokerResponse, error)
}

type idPOSRepository struct {
	db *pinot.Connection
}

func NewIdPOSRepository(db *pinot.Connection) *idPOSRepository {
	return &idPOSRepository{
		db: db,
	}
}

func (r *idPOSRepository) GetIdPOS(ctx context.Context, limit int) (*pinot.BrokerResponse, error) {
	table := "idpos"
	query := fmt.Sprintf("select * from idpos limit %d", limit)

	brokerResp, err := r.db.ExecuteSQL(table, query)
	if err != nil {
		log.Error(err)
		return &pinot.BrokerResponse{}, err
	}

	return brokerResp, nil
}
