package item

type StocksProvider interface {
	GetStocks(sku uint32) (uint64, error)
}

type ProductInfoProvider interface {
	GetProductInfo(sku uint32) (string, uint32, error)
}

type AddService struct {
	stocksProvider      StocksProvider
	productInfoProvider ProductInfoProvider
}

func NewAddService(stocksProvider StocksProvider, productInfoProvider ProductInfoProvider) *AddService {

	return &AddService{
		stocksProvider:      stocksProvider,
		productInfoProvider: productInfoProvider,
	}
}

func (s AddService) Add(User int64, SKU uint32, Count uint16) error {
	return nil
}
