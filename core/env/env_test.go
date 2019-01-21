package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsProduction(t *testing.T) {
	t.Run("環境変数のMODEが設定されてない場合はFalseを戻す", func(t *testing.T) {
		assert.False(t, IsProduction())
	})

	t.Run("環境変数のMODEがproductionの場合はTrueを戻す", func(t *testing.T) {
		os.Setenv("MODE", "production")
		assert.True(t, IsProduction())
	})
}
