package orders

import "errors"

var (
	ErrOrderIDIncorrect = errors.New("order ID incorrect")
	ErrUserIncorrect    = errors.New("user incorrect")

	ErrItemsSKUIncorrect   = errors.New("Items.%d SKU incorrect")
	ErrItemsCountIncorrect = errors.New("Items.%d Count incorrect")
)
