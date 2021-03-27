package processData

func DailyTotal(input map[string]interface{}) map[string]float64 {
	dailyTotal := make(map[string]float64)
	for k, v := range input {
		switch vv := v.(type) {
		case map[string]interface{}:
			grandTotal := vv["grand_total"].(map[string]interface{})
			grandTotalInSeconds := grandTotal["total_seconds"].(float64)
			dailyTotal[k] = grandTotalInSeconds
		default:
			// Do nothing
		}
	}
	return dailyTotal
}
