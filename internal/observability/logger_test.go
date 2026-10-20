package observability

import (
	"os"
	"testing"
)

func TestNewLogger(t *testing.T) {
	dir := t.TempDir()
	ch, err := NewLogger(dir, "info", 1, 1, 1, "test")
	if err != nil {
		t.Fatal(err)
	}
	if ch.App == nil || ch.Error == nil || ch.Audit == nil {
		t.Fatal("nil loggers")
	}
	ch.App.Info("hello")
	// ensure file created
	entries, _ := os.ReadDir(dir)
	if len(entries) == 0 {
		t.Fatal("no log files")
	}
}