package goutils_test

import (
	"regexp"
	"testing"

	"github.com/mekramy/goutils"
)

func TestExtractNumbers(t *testing.T) {
	input := "abc123def456"
	expected := "123456"
	result := goutils.ExtractNumbers(input)
	if result != expected {
		t.Errorf("ExtractNumbers(%q) = %q; want %q", input, result, expected)
	}
}

func TestExtractAlphaNum(t *testing.T) {
	input := "abc123!@_def456"
	expected := "abc123def456"
	result := goutils.ExtractAlphaNum(input)
	if result != expected {
		t.Errorf("ExtractAlphaNum(%q) = %q; want %q", input, result, expected)
	}
}

func TestExtractAlphaNumPersian(t *testing.T) {
	input := "abc123!@_def456سلام گچپژ"
	expected := "abc123def456سلامگچپژ"
	result := goutils.ExtractAlphaNumPersian(input)
	if result != expected {
		t.Errorf("ExtractAlphaNumPersian(%q) = %q; want %q", input, result, expected)
	}
}

func TestRandomNumeric(t *testing.T) {
	length := uint(10)
	result := goutils.RandomNumeric(length)
	if len(result) != int(length) {
		t.Errorf("RandomNumeric(%d) = %q; length = %d; want %d", length, result, len(result), length)
	}
	if match, _ := regexp.MatchString(`^\d+$`, result); !match {
		t.Errorf("RandomNumeric(%d) = %q; want numeric string", length, result)
	}
}

func TestRandomAlphaNum(t *testing.T) {
	length := uint(10)
	result := goutils.RandomAlphaNum(length)
	if len(result) != int(length) {
		t.Errorf("RandomAlphaNum(%d) = %q; length = %d; want %d", length, result, len(result), length)
	}
	if match, _ := regexp.MatchString(`^[A-Z0-9]+$`, result); !match {
		t.Errorf("RandomAlphaNum(%d) = %q; want alphanumeric string", length, result)
	}
}

func TestSlugify(t *testing.T) {
	input := []string{"Hello-- ", "  World!"}
	expected := "Hello-World"
	result := goutils.Slugify(input...)
	if result != expected {
		t.Errorf("Slugify(%q) = %q; want %q", input, result, expected)
	}
}

func TestSlugifyPersian(t *testing.T) {
	input := []string{"سلام", "دنیا!"}
	expected := "سلام-دنیا"
	result := goutils.SlugifyPersian(input...)
	if result != expected {
		t.Errorf("SlugifyPersian(%q) = %q; want %q", input, result, expected)
	}
}

func TestConcat(t *testing.T) {
	input := []string{"Hello", "", "      ", "World"}
	sep := " "
	expected := "Hello World"
	result := goutils.Concat(sep, input...)
	if result != expected {
		t.Errorf("Concat(%q, %q) = %q; want %q", sep, input, result, expected)
	}
}

func TestFormatNumber(t *testing.T) {
	layout := "%d Dollars"
	value := 100000
	expected := "100,000 Dollars"
	result := goutils.FormatNumber(layout, value)
	if result != expected {
		t.Errorf("FormatNumber(%q, %d) = %q; want %q", layout, value, result, expected)
	}
}

func TestFormatRx(t *testing.T) {
	data := "123456"
	pattern := `(\d{3})(\d{2})(\d{1})`
	repl := "($1) $2-$3"
	expected := "(123) 45-6"
	result, err := goutils.FormatRx(data, pattern, repl)
	if err != nil {
		t.Errorf("FormatRx(%q, %q, %q) returned error: %v", data, pattern, repl, err)
	}
	if result != expected {
		t.Errorf("FormatRx(%q, %q, %q) = %q; want %q", data, pattern, repl, result, expected)
	}
}
