package e2e

import (
	"os"
	"testing"

	zone "github.com/lrstanley/bubblezone"
)

func TestMain(m *testing.M) {
	zone.NewGlobal()
	zone.SetEnabled(false)
	os.Exit(m.Run())
}
