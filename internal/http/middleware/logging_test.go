package middleware

import "testing"

func TestRedact(t *testing.T) {
	in := `{"password":"secret","token":"abc","api_key":"k","authorization":"Bearer x"}`
	out := redact(in)
	if out == in {
		t.Fatalf("expected redaction, got %s", out)
	}
	if contains(out, "secret") || contains(out, "abc") {
		t.Fatalf("sensitive value not redacted: %s", out)
	}
}

func contains(s, sub string) bool {
	return len(s) >= len(sub) && (func() bool {
		for i := 0; i+len(sub) <= len(s); i++ {
			if s[i:i+len(sub)] == sub {
				return true
			}
		}
		return false
	})()
}