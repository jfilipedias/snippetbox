package assert

import (
	"html"
	"regexp"
	"strings"
	"testing"
)

func Equal[T comparable](t *testing.T, got, want T) {
	t.Helper()

	if got != want {
		t.Errorf("got: %v; want: %v", got, want)
	}
}

func StringContains(t *testing.T, got, want string) {
	t.Helper()

	if !strings.Contains(got, want) {
		t.Errorf("got: %q; expected to contain: %q", got, want)
	}
}

var csrfTokenRx = regexp.MustCompile(`<input type='hidden' name='csrf_token' value='(.+)'>`)

func ExtractCSRFToken(t *testing.T, body string) string {
	matches := csrfTokenRx.FindStringSubmatch(body)
	if len(matches) < 2 {
		t.Fatal("no csrf token found in body")
	}

	return html.UnescapeString(matches[1])
}
