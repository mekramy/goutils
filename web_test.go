package goutils_test

import (
	"testing"

	"github.com/mekramy/goutils"
)

func TestRelativeURL(t *testing.T) {
	tests := []struct {
		root     string
		path     []string
		expected string
	}{
		{"g:/mekramy", []string{"g:/mekramy/goutils/web.go"}, "goutils/web.go"},
		{"g:/mekramy", []string{"g:/mekramy/goutils"}, "goutils"},
		{"g:/mekramy", []string{"g:/mekramy"}, ""},
	}

	for _, test := range tests {
		result := goutils.RelativeURL(test.root, test.path...)
		if result != test.expected {
			t.Errorf("RelativeURL(%q, %q) = %q; want %q", test.root, test.path, result, test.expected)
		}
	}
}

func TestAbsoluteURL(t *testing.T) {
	tests := []struct {
		root     string
		path     []string
		expected string
	}{
		{"g:/mekramy", []string{"g:/mekramy/goutils/web.go"}, "/goutils/web.go"},
		{"g:/mekramy", []string{"g:/mekramy/goutils"}, "/goutils"},
		{"g:/mekramy", []string{"g:/mekramy"}, "/"},
	}

	for _, test := range tests {
		result := goutils.AbsoluteURL(test.root, test.path...)
		if result != test.expected {
			t.Errorf("AbsoluteURL(%q, %q) = %q; want %q", test.root, test.path, result, test.expected)
		}
	}
}

func TestSanitizeRaw(t *testing.T) {
	tests := []struct {
		data     string
		trim     bool
		expected string
	}{
		{"<script>alert('xss')</script>", false, ""},
		{"<b>bold</b>", true, "bold"},
		{"  <i>italic</i>  ", true, "italic"},
	}

	for _, test := range tests {
		result := goutils.SanitizeRaw(test.data, test.trim)
		if result != test.expected {
			t.Errorf("SanitizeRaw(%q, %t) = %q; want %q", test.data, test.trim, result, test.expected)
		}
	}
}

func TestSanitizeCommon(t *testing.T) {
	tests := []struct {
		data     string
		trim     bool
		expected string
	}{
		{"<script>alert('xss')</script>", false, ""},
		{"<b>bold</b>", true, "<b>bold</b>"},
		{"  <i>italic</i>  ", true, "<i>italic</i>"},
	}

	for _, test := range tests {
		result := goutils.SanitizeCommon(test.data, test.trim)
		if result != test.expected {
			t.Errorf("SanitizeCommon(%q, %t) = %q; want %q", test.data, test.trim, result, test.expected)
		}
	}
}
