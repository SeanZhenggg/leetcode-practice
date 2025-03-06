package binarySearch

import "fmt"

type kv struct {
	Ts    int
	Value string
}
type TimeMap struct {
	m map[string][]kv
}

func TimeBasedKVStoreConstructor() TimeMap {
	return TimeMap{
		m: make(map[string][]kv),
	}
}

func (tm *TimeMap) Set(key string, value string, timestamp int) {
	if _, found := tm.m[key]; !found {
		tm.m[key] = make([]kv, 0)
	}

	tm.m[key] = append(tm.m[key], kv{timestamp, value})
}

func (tm *TimeMap) Get(key string, timestamp int) string {
	list := tm.m[key]

	l, r := 0, len(list)-1

	for l <= r {
		mid := l + (r-l)/2

		if list[mid].Ts < timestamp {
			l = mid + 1
		} else if list[mid].Ts == timestamp {
			return list[mid].Value
		} else {
			r = mid - 1
		}
	}

	if r != -1 {
		return list[r].Value
	}

	return ""
}

type TimeMap2 struct {
	m map[string][]kv
}

func TimeBasedKVStoreConstructorReview() TimeMap {
	return TimeMap{
		m: make(map[string][]kv),
	}
}

func (tm *TimeMap2) Set(key string, value string, timestamp int) {
	if _, found := tm.m[key]; !found {
		tm.m[key] = make([]kv, 0)
	}

	tm.m[key] = append(tm.m[key], kv{timestamp, value})
}

func (tm *TimeMap2) Get(key string, timestamp int) string {
	list := tm.m[key]

	l, r := 0, len(list)-1

	for l <= r {
		mid := l + (r-l)/2

		if list[mid].Ts <= timestamp {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	if l == 0 {
		return ""
	}

	return list[l-1].Value
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */

func Test_TimeBasedKeyValueStore() {
	ans := make([]*string, 0, 7)
	obj := TimeBasedKVStoreConstructor()
	ans = append(ans, nil)
	obj.Set("foo", "bar", 1)
	ans = append(ans, nil)
	val1 := obj.Get("foo", 1)
	ans = append(ans, &val1)
	val2 := obj.Get("foo", 3)
	ans = append(ans, &val2)
	obj.Set("foo", "bar2", 4)
	ans = append(ans, nil)
	val3 := obj.Get("foo", 4)
	ans = append(ans, &val3)
	val4 := obj.Get("foo", 5)
	ans = append(ans, &val4)
	val5 := obj.Get("foo", 3)
	ans = append(ans, &val5)

	for _, answer := range ans {
		if answer == nil {
			fmt.Printf("%s\t", "null")
		} else {
			fmt.Printf("%s\t", *answer)
		}
	}
}

func Test_TimeBasedKeyValueStoreReview() {
	ans := make([]*string, 0, 7)
	obj := TimeBasedKVStoreConstructorReview()
	ans = append(ans, nil)
	obj.Set("foo", "bar", 1)
	ans = append(ans, nil)
	val1 := obj.Get("foo", 1)
	ans = append(ans, &val1)
	val2 := obj.Get("foo", 3)
	ans = append(ans, &val2)
	obj.Set("foo", "bar2", 4)
	ans = append(ans, nil)
	val3 := obj.Get("foo", 4)
	ans = append(ans, &val3)
	val4 := obj.Get("foo", 5)
	ans = append(ans, &val4)
	val5 := obj.Get("foo", 3)
	ans = append(ans, &val5)

	for _, answer := range ans {
		if answer == nil {
			fmt.Printf("%s\t", "null")
		} else {
			fmt.Printf("%s\t", *answer)
		}
	}
}
