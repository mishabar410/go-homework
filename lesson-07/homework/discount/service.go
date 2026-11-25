package discount

import "errors"

var (
	ErrNegativePrice  = errors.New("price must not be negative")
	ErrInvalidPercent = errors.New("discount percent must be between 0 and 100")
)

type Item struct {
	Price           int
	DiscountPercent int
}

func ValidatePercent(percent int) error {
	if percent < 0 || percent > 100 {
		return ErrInvalidPercent
	}
	return nil
}

func Apply(price, percent int) (int, error) {
	if price < 0 {
		return 0, ErrNegativePrice
	}
	if err := ValidatePercent(percent); err != nil {
		return 0, err
	}

	return price * (100 - percent) / 100, nil
}

func FinalPrice(items []Item) (int, error) {
	total := 0

	for _, item := range items {
		price, err := Apply(item.Price, item.DiscountPercent)
		if err != nil {
			return 0, err
		}
		total += price
	}

	return total, nil
}
