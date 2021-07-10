package main

// 981 https://leetcode-cn.com/problems/time-based-key-value-store/
// 哈希 + 二分查找
type TimeMap struct {
	m map[string][]item
}

type item struct {
	value     string
	timestamp int
}

/** Initialize your data structure here. */
func Constructor() TimeMap {
	this := TimeMap{
		m: make(map[string][]item),
	}

	return this
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if _, ok := this.m[key]; !ok {
		this.m[key] = make([]item, 0)
	}

	this.m[key] = append(this.m[key], item{
		value:     value,
		timestamp: timestamp,
	})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	items := this.m[key]
	if items == nil {
		return ""
	}

	start := 0
	end := len(items) - 1

	index := 0
	for start <= end {
		mid := (start + end) / 2
		// 如果合法则更新 index
		if this.m[key][mid].timestamp <= timestamp {
			index = mid
		}

		if this.m[key][mid].timestamp < timestamp {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	// 判断是否合法
	if this.m[key][index].timestamp > timestamp {
		return ""
	}

	return this.m[key][index].value
}
