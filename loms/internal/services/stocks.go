package services

type StocksProvider interface {
	GetStocks(sku uint32) uint64
}

type StocksService struct {
	stocksProvider StocksProvider
}

func NewStocksService(stocksProvider StocksProvider) *StocksService {

	return &StocksService{
		stocksProvider: stocksProvider,
	}
}

func (s StocksService) GetStocks(sku uint32) uint64 {
	return s.stocksProvider.GetStocks(sku)
}
