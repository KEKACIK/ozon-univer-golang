package repository

import (
	"fmt"
	"net/http"
	"route256/loms/internal/models"
	"sync"
)

type DumpRepo struct {
	mu sync.Mutex

	orders map[int64]models.OrderModel
	stocks map[uint32]models.StockModel
}

func NewDumpRepo() *DumpRepo {
	orders := map[int64]models.OrderModel{}
	stocks := map[uint32]models.StockModel{
		uint32(1076963): {SKU: 1076963, Count: 15, Reserved: 2},
		uint32(5678901): {SKU: 5678901, Count: 62, Reserved: 0},
		uint32(4200000): {SKU: 4200000, Count: 543, Reserved: 18},
		uint32(2500001): {SKU: 2500001, Count: 900, Reserved: 0},
		uint32(8765432): {SKU: 8765432, Count: 122, Reserved: 0},
		uint32(7000123): {SKU: 7000123, Count: 1, Reserved: 0},
	}

	return &DumpRepo{
		orders: orders,
		stocks: stocks,
	}
}

func (dr *DumpRepo) Test(w http.ResponseWriter, r *http.Request) {
	fmt.Println(dr.stocks)
	fmt.Println(dr.orders)
}
