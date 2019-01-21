package config

import (
	"testing"
)

func TestNewConfig(t *testing.T) {
	t.Run("設定ファイルがない場合は読み込まないこと", func(t *testing.T) {
		// got := NewConfig(l)
		// assert.False(t, got.InConfig("config/app.yaml"))
	})

	t.Run("設定ファイルがある場合は読み込む", func(t *testing.T) {
		// got := NewConfig(l)
		// assert.False(t, got.InConfig("config/app.yaml"))
	})
}
