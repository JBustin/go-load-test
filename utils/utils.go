package utils

import (
	"math/rand"
	"time"
)

func Fill(urls []string, hits int) []string {
	if len(urls) == 0 {
		return urls
	}

	if len(urls) == 0 && len(urls) >= hits {
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

func Shuffle(mySlice []string) []string {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(mySlice), func(i, j int) { mySlice[i], mySlice[j] = mySlice[j], mySlice[i] })
	return mySlice
}
