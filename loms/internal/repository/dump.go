package repository

import "math/rand"

type DumpRepo struct {
}

func NewDumpRepo() *DumpRepo {

	return &DumpRepo{}
}

func (r DumpRepo) GetStocks(sku uint32) uint64 {
	return uint64(rand.Int() % 1000)
}
