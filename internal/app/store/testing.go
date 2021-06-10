package store

import (
	"fmt"
	"strings"
	"testing"

	"github.com/lehius/http-resy-api/internal/config"
)

// TestStore
func TestStore(t *testing.T, databaseURL string) (*Store, func(...string)) {
	t.Helper()

	config := config.Default()
	s := New(config.Store)
	err := s.Open()
	if err != nil {
		t.Fatal(err)
	}

	return s, func(tables ...string) {
		if len(tables) > 0 {
			if _, err = s.db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))); err != nil {
				t.Fatal(err)
			}
		}
		s.Close()
	}
}
