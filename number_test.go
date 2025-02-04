package goutils_test

import (
	"testing"

	"github.com/mekramy/goutils"
)

func TestAbs(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected interface{}
	}{
		{input: int(-5), expected: int(5)},
		{input: int(5), expected: int(5)},
		{input: float64(-5.5), expected: float64(5.5)},
		{input: float64(5.5), expected: float64(5.5)},
	}

	for _, test := range tests {
		switch v := test.input.(type) {
		case int:
			if result := goutils.Abs(v); result != test.expected {
				t.Errorf("Abs(%v) = %v; want %v", v, result, test.expected)
			}
		case float64:
			if result := goutils.Abs(v); result != test.expected {
				t.Errorf("Abs(%v) = %v; want %v", v, result, test.expected)
			}
		}
	}
}

func TestRoundUp(t *testing.T) {
	tests := []struct {
		input    float64
		expected int
	}{
		{input: 5.1, expected: 6},
		{input: 5.9, expected: 6},
		{input: -5.1, expected: -5},
		{input: -5.9, expected: -5},
	}

	for _, test := range tests {
		if result := goutils.RoundUp[int](test.input); result != test.expected {
			t.Errorf("RoundUp(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestRound(t *testing.T) {
	tests := []struct {
		input    float64
		expected int
	}{
		{input: 5.5, expected: 6},
		{input: 5.4, expected: 5},
		{input: -5.5, expected: -6},
		{input: -5.4, expected: -5},
	}

	for _, test := range tests {
		if result := goutils.Round[int](test.input); result != test.expected {
			t.Errorf("Round(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestRoundDown(t *testing.T) {
	tests := []struct {
		input    float64
		expected int
	}{
		{input: 5.9, expected: 5},
		{input: 5.1, expected: 5},
		{input: -5.9, expected: -5},
		{input: -5.1, expected: -5},
	}

	for _, test := range tests {
		if result := goutils.RoundDown[int](test.input); result != test.expected {
			t.Errorf("RoundDown(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		input    []interface{}
		expected interface{}
	}{
		{input: []interface{}{int(5), int(3), int(9)}, expected: int(3)},
		{input: []interface{}{float64(5.5), float64(3.3), float64(9.9)}, expected: float64(3.3)},
	}

	for _, test := range tests {
		switch test.input[0].(type) {
		case int:
			if result := goutils.Min(test.input[0].(int), test.input[1].(int), test.input[2].(int)); result != test.expected {
				t.Errorf("Min(%v) = %v; want %v", test.input, result, test.expected)
			}
		case float64:
			if result := goutils.Min(test.input[0].(float64), test.input[1].(float64), test.input[2].(float64)); result != test.expected {
				t.Errorf("Min(%v) = %v; want %v", test.input, result, test.expected)
			}
		}
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		input    []interface{}
		expected interface{}
	}{
		{input: []interface{}{int(5), int(3), int(9)}, expected: int(9)},
		{input: []interface{}{float64(5.5), float64(3.3), float64(9.9)}, expected: float64(9.9)},
	}

	for _, test := range tests {
		switch test.input[0].(type) {
		case int:
			if result := goutils.Max(test.input[0].(int), test.input[1].(int), test.input[2].(int)); result != test.expected {
				t.Errorf("Max(%v) = %v; want %v", test.input, result, test.expected)
			}
		case float64:
			if result := goutils.Max(test.input[0].(float64), test.input[1].(float64), test.input[2].(float64)); result != test.expected {
				t.Errorf("Max(%v) = %v; want %v", test.input, result, test.expected)
			}
		}
	}
}
