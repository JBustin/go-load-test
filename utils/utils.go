package utils

func Fill(urls []string, hits int) []string {
	if len(urls) == 0 || len(urls) >= hits {
		return urls[0:hits]
	}

	result := make([]string, hits)

	var j int
	for i := 0; i < hits; i++ {
		if j >= len(urls) {
			j = 0
		}
		result[i] = urls[j]
		j++
	}

	return result
}

func Dedupe(mySlice []string) []string {
	result := []string{}
	deduper := make(map[string]bool)

	for _, value := range mySlice {
		_, exist := deduper[value]
		if !exist {
			result = append(result, value)
			deduper[value] = true
		}
	}

	return result
}
