package repository

import (
	"errors"
)

var (
	ErrStockNotFound  = errors.New("stock not found")
	ErrStockNotEnough = errors.New("stock not enough")
)

func (r *DumpRepo) GetStocks(sku uint32) (uint64, error) {
	stock, ok := r.stocks[sku]
	if !ok {
		return 0, ErrStockNotFound
	}

	return uint64(stock.Count), nil
}

func (r *DumpRepo) ReserveStock(sku uint32, count uint16) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	stock, ok := r.stocks[sku]
	if !ok {
		return ErrStockNotFound
	}
	freeCount := stock.Count - stock.Reserved
	if freeCount < count {
		return ErrStockNotEnough
	}

	stock.Reserved += count
	r.stocks[sku] = stock

	return nil
}

func (r *DumpRepo) UnReserveStock(sku uint32, count uint16) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	stock, ok := r.stocks[sku]
	if !ok {
		return ErrStockNotFound
	}
	if stock.Reserved > count {
		return ErrStockNotEnough
	}

	stock.Reserved -= count
	r.stocks[sku] = stock

	return nil
}
