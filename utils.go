package ghealth

func contains(items []HealtCheckResponseItem, condition string) bool {
	for _, v := range items {
		if v.Status == condition {
			return true
		}
	}

	return false
}
