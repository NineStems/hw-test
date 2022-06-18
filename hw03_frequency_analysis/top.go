package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type counter struct {
	count int
	value string
}

func Top10(row string) []string {
	if row == "" {
		return nil
	}
	uniqCountMap := getUniqMap(strings.Fields(row))
	return transformToSortSlice(uniqCountMap)
}

func getUniqMap(words []string) map[string]counter {
	if len(words) == 0 {
		return nil
	}
	uniq := map[string]counter{}
	for _, word := range words {
		obj, has := uniq[word]
		if !has {
			obj = counter{value: word}
		}
		obj.count++
		uniq[word] = obj
	}
	return uniq
}

func transformToSortSlice(uniq map[string]counter) []string {
	if len(uniq) == 0 {
		return nil
	}
	cntArr := make([]counter, 0, len(uniq))
	for _, value := range uniq {
		cntArr = append(cntArr, value)
	}
	sort.Slice(cntArr, func(i, j int) bool {
		if cntArr[i].count != cntArr[j].count {
			return cntArr[i].count > cntArr[j].count
		}
		return cntArr[i].value < cntArr[j].value
	})
	var res []string
	switch {
	case len(cntArr) > 10:
		res = make([]string, 10)
	default:
		res = make([]string, len(cntArr))
	}
	for i := range res {
		res[i] = cntArr[i].value
	}
	return res
}
