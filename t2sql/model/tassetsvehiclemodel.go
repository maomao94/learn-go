package model

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TAssetsVehicleModel = (*customTAssetsVehicleModel)(nil)

type (
	// TAssetsVehicleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTAssetsVehicleModel.
	TAssetsVehicleModel interface {
		tAssetsVehicleModel
	}

	customTAssetsVehicleModel struct {
		*defaultTAssetsVehicleModel
	}
)

func (c customTAssetsVehicleModel) FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*TAssetsVehicle, error) {
	if orderBy == "" {
		rowBuilder = rowBuilder.OrderBy("id DESC")
	} else {
		rowBuilder = rowBuilder.OrderBy(orderBy)
	}
	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}
	var resp []*TAssetsVehicle
	err = c.conn.QueryRowsCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

// NewTAssetsVehicleModel returns a model for the database table.
func NewTAssetsVehicleModel(conn sqlx.SqlConn) TAssetsVehicleModel {
	return &customTAssetsVehicleModel{
		defaultTAssetsVehicleModel: newTAssetsVehicleModel(conn),
	}
}
