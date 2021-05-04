package mapslice

import "sort"

// KsVs2Slc : orderType [K-DESC / K-ASC / V-DESC / V-ASC] by string length
// map[string]string => key []string & value []string
func KsVs2Slc(m map[string]string, orderType string) (keys []string, values []string) {

	if m == nil {
		return nil, nil
	}
	if len(m) == 0 {
		return []string{}, []string{}
	}

	type kv struct {
		key   string
		value string
	}

	kvSlc := []kv{}
	for k, v := range m {
		kvSlc = append(kvSlc, kv{key: k, value: v})
	}

	switch orderType {
	case "K-DESC":
		sort.SliceStable(kvSlc, func(i, j int) bool { return len(kvSlc[i].key) > len(kvSlc[j].key) })
	case "K-ASC":
		sort.SliceStable(kvSlc, func(i, j int) bool { return len(kvSlc[i].key) < len(kvSlc[j].key) })
	case "V-DESC":
		sort.SliceStable(kvSlc, func(i, j int) bool { return len(kvSlc[i].value) > len(kvSlc[j].value) })
	case "V-ASC":
		sort.SliceStable(kvSlc, func(i, j int) bool { return len(kvSlc[i].value) < len(kvSlc[j].value) })
	}

	for _, kvEle := range kvSlc {
		keys = append(keys, kvEle.key)
		values = append(values, kvEle.value)
	}
	return
}
