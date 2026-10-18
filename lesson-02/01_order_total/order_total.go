package order_total

func CalculateOrder(prices []int64, quantities []int, discount int) (subtotal, discountAmount, total int64) {
	if len(prices) != len(quantities) || discount < 0 || discount > 100 {
		return 0, 0, 0
	}

	for i, price := range prices {
		quantity := quantities[i]
		if price < 0 || quantity < 0 {
			return 0, 0, 0
		}
		subtotal += price * int64(quantity)
	}

	d := int64(discount)
	discountAmount = subtotal/100*d + subtotal%100*d/100
	total = subtotal - discountAmount

	return subtotal, discountAmount, total
}
