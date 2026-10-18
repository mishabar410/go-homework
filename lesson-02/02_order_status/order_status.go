package order_status

func CountStatuses(statuses []string) (newCount, paidCount, shippedCount, cancelledCount, unknownCount, activeCount int) {
	for _, status := range statuses {
		switch status {
		case "new":
			newCount++
		case "paid":
			paidCount++
		case "shipped":
			shippedCount++
		case "cancelled":
			cancelledCount++
		default:
			unknownCount++
		}
	}

	activeCount = newCount + paidCount + shippedCount

	return newCount, paidCount, shippedCount, cancelledCount, unknownCount, activeCount
}
