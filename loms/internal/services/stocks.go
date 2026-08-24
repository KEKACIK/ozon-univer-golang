package services

import (
	"context"

	"github.com/KEKACIK/ozon-univer-golang/loms/internal/repository/stocks"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StocksService struct {
	dbPool *pgxpool.Pool
}

func NewStocksService(dbPool *pgxpool.Pool) *StocksService {

	return &StocksService{dbPool: dbPool}
}

func (s *StocksService) GetStocks(ctx context.Context, sku uint32) (uint64, error) {
	stocksRepo := stocks.New(s.dbPool)

	stock, err := stocksRepo.GetStocks(context.Background(), int32(sku))
	if err != nil {
		return 0, err
	}

	return uint64(stock.Count) - uint64(stock.Reserved), nil
}
