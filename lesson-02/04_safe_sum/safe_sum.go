package safe_sum

import "math"

func SafeSum(transactions []int64) (total int64, skipped []int) {
	for i, t := range transactions {
		overflow := (t > 0 && total > math.MaxInt64-t) ||
			(t < 0 && total < math.MinInt64-t)

		if overflow {
			skipped = append(skipped, i)
			continue
		}

		total += t
	}

	return total, skipped
}
