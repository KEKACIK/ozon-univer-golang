package item

import "errors"

type StocksProvider interface {
	GetStocks(sku uint32) (uint64, error)
}

type ProductProvider interface {
	GetProductInfo(sku uint32) (string, uint32, error)
}

type AddService struct {
	stocksProvider  StocksProvider
	productProvider ProductProvider
}

func NewAddService(stocksProvider StocksProvider, productProvider ProductProvider) *AddService {

	return &AddService{
		stocksProvider:  stocksProvider,
		productProvider: productProvider,
	}
}

var (
	ErrIncorrectCount = errors.New("incorrect quantity")
)

func (s AddService) Add(User int64, SKU uint32, Count uint16) error {
	if _, _, err := s.productProvider.GetProductInfo(SKU); err != nil {
		return err
	}

	stockCount := uint16(1000)
	if stockCount < Count {
		return ErrIncorrectCount
	}

	return nil
}
