package stack

import (
	"reflect"
	"testing"
)

type DailyTemperaturesCase struct {
	input    []int
	expected []int
}

var dailyTemperaturesCases = []DailyTemperaturesCase{
	{input: []int{73, 74, 75, 71, 69, 72, 76, 73}, expected: []int{1, 1, 4, 2, 1, 1, 0, 0}},
	{input: []int{30, 40, 50, 60}, expected: []int{1, 1, 1, 0}},
	{input: []int{30, 60, 90}, expected: []int{1, 1, 0}},
}

func TestDailyTemperatures(t *testing.T) {
	for _, c := range dailyTemperaturesCases {
		ans := dailyTemperatures(c.input)
		if !reflect.DeepEqual(ans, c.expected) {
			t.Errorf("answer is %v, want %v", ans, c.expected)
		}
	}
}

func TestDailyTemperatures2(t *testing.T) {
	for _, c := range dailyTemperaturesCases {
		ans := dailyTemperatures2(c.input)
		if !reflect.DeepEqual(ans, c.expected) {
			t.Errorf("answer is %v, want %v", ans, c.expected)
		}
	}
}

func TestDailyTemperatures3(t *testing.T) {
	for _, c := range dailyTemperaturesCases {
		ans := dailyTemperatures3(c.input)
		if !reflect.DeepEqual(ans, c.expected) {
			t.Errorf("answer is %v, want %v", ans, c.expected)
		}
	}
}

func TestDailyTemperaturesReview(t *testing.T) {
	for _, c := range dailyTemperaturesCases {
		ans := dailyTemperaturesReview(c.input)
		if !reflect.DeepEqual(ans, c.expected) {
			t.Errorf("answer is %v, want %v", ans, c.expected)
		}
	}
}

func TestDailyTemperaturesReview2(t *testing.T) {
	for _, c := range dailyTemperaturesCases {
		ans := dailyTemperaturesReview2(c.input)
		if !reflect.DeepEqual(ans, c.expected) {
			t.Errorf("answer is %v, want %v", ans, c.expected)
		}
	}
}

func TestDailyTemperaturesMonotonicStack(t *testing.T) {
	for _, c := range dailyTemperaturesCases {
		ans := dailyTemperaturesMonotonicStack(c.input)
		if !reflect.DeepEqual(ans, c.expected) {
			t.Errorf("answer is %v, want %v", ans, c.expected)
		}
	}
}

func TestDailyTemperaturesArray(t *testing.T) {
	for _, c := range dailyTemperaturesCases {
		ans := dailyTemperaturesArray(c.input)
		if !reflect.DeepEqual(ans, c.expected) {
			t.Errorf("answer is %v, want %v", ans, c.expected)
		}
	}
}
