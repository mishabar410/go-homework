package order_retry

func ProcessOrder(results []bool, maxAttempts int) (success bool, attempts int) {
	for _, value := range results {
		if attempts >= maxAttempts {
			break
		}

		attempts = attempts + 1

		if value {
			return true, attempts
		}
	}
	return false, attempts
}
