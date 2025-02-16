package goutils

import (
	"html"
	"path/filepath"
	"strings"

	"github.com/microcosm-cc/bluemonday"
)

// RelativeURL returns the relative URL path of
// file with respect to the root directory.
func RelativeURL(root string, path ...string) string {
	normalizedRoot := NormalizePath(root)
	normalizedPath := NormalizePath(path...)
	relPath := strings.TrimPrefix(normalizedPath, normalizedRoot)
	relPath = strings.TrimPrefix(relPath, string(filepath.Separator))
	return filepath.Clean(filepath.ToSlash(relPath))
}

// AbsoluteURL returns the absolute URL path of
// file with respect to the root directory.
func AbsoluteURL(root string, path ...string) string {
	return filepath.Clean("/" + RelativeURL(root, path...))
}

// SanitizeRaw sanitize input to raw text.
func SanitizeRaw(data string, trim bool) string {
	clean := bluemonday.StrictPolicy().Sanitize(data)
	if trim {
		clean = strings.TrimSpace(clean)
	}
	return html.UnescapeString(clean)
}

// SanitizeCommon sanitize input to html with common allowed tags.
func SanitizeCommon(data string, trim bool) string {
	clean := bluemonday.UGCPolicy().Sanitize(data)
	if trim {
		clean = strings.TrimSpace(clean)
	}
	return html.UnescapeString(clean)
}
