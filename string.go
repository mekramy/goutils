package goutils

import (
	"regexp"
	"strings"
	"time"

	"golang.org/x/exp/rand"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// ExtractNumbers extract numbers from string.
func ExtractNumbers(s string) string {
	rx := regexp.MustCompile(`[^0-9]`)
	return rx.ReplaceAllString(s, "")
}

// ExtractAlphaNum extract alpha and numbers from string [a-zA-Z0-9].
func ExtractAlphaNum(s string, includes ...string) string {
	pattern := "[^a-zA-Z0-9" + strings.Join(includes, "") + "]"
	rx := regexp.MustCompile(pattern)
	return rx.ReplaceAllString(s, "")
}

// ExtractAlphaNumPersian extract english and persian alpha and numbers from string [ا-یa-zA-Z0-9].
func ExtractAlphaNumPersian(s string, includes ...string) string {
	pattern := "[^\u0600-\u06FF\uFB8A\u067E\u0686\u06AFa-zA-Z0-9" + strings.Join(includes, "") + "]"
	rx := regexp.MustCompile(pattern)
	return rx.ReplaceAllString(s, "")
}

// RandomNumeric returns random string from character set.
func RandomString(n uint, characters string) string {
	src := rand.NewSource(uint64(time.Now().UnixNano()))
	rnd := rand.New(src)

	result := make([]byte, n)
	for i := range result {
		result[i] = characters[rnd.Intn(len(characters))]
	}
	return string(result)
}

// RandomNumeric returns random numeric string.
func RandomNumeric(n uint) string {
	return RandomString(n, "0123456789")
}

// RandomAlphaNum returns random string from Alpha-Num uppercase characters.
func RandomAlphaNum(n uint) string {
	return RandomString(n, "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
}

// Slugify make url friendly slug from strings.
// Only Alpha-Num characters compiled to final result.
func Slugify(parts ...string) string {
	normalized := ExtractAlphaNum(strings.Join(parts, " "), `\s\-`)
	rx := regexp.MustCompile(`[\s\-]+`)
	return rx.ReplaceAllString(normalized, "-")
}

// SlugifyPersian make url friendly slug from strings.
// Only english and persian alpha and numberic characters compiled to final result.
func SlugifyPersian(parts ...string) string {
	normalized := ExtractAlphaNumPersian(strings.Join(parts, " "), `\s\-`)
	rx := regexp.MustCompile(`[\s\-]+`)
	return rx.ReplaceAllString(normalized, "-")
}

// Concat return concatinated not-empty strings with separator.
func Concat(sep string, parts ...string) string {
	res := make([]string, 0)
	for _, part := range parts {
		if strings.TrimSpace(part) != "" {
			res = append(res, part)
		}
	}
	return strings.Join(res, sep)
}

// FormatNumber format number with comma separator.
//
// code block:
//
//	FormatNumber("%d Dollars", 100000)
//
// output:
//
//	100,000 Dollars
func FormatNumber(layout string, v ...any) string {
	p := message.NewPrinter(language.English)
	return p.Sprintf(layout, v...)
}

// FormatRx format string using regex pattern
// use () for match groups and $1, $2 for output placeholder.
//
// code block:
//
//	FormatRx("123456", `(\d{3})(\d{2})(\d{1})`, "($1) $2-$3")
//
// output:
//
//	(123) 45-6
func FormatRx(data, pattern, repl string) (string, error) {
	rx, err := regexp.Compile("^" + pattern + "$")
	if err != nil {
		return "", err
	}
	return rx.ReplaceAllString(data, repl), nil
}
