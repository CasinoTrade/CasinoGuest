package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/CasinoTrade/CasinoGuest/internal/model/config"
)

func TestTomlMarshaling(t *testing.T) {
	t.Parallel()

	t.Run("marshal default config", func(t *testing.T) {
		t.Parallel()

		cfg := config.DefaultCfg()
		assert.NoError(t, config.Validate(cfg))

		data, err := config.MarshalToml(cfg)
		assert.NoError(t, err)
		assert.NotNil(t, data)

		cfg2, err := config.UnmarshalToml(data)
		assert.NoError(t, err)

		assert.Equal(t, cfg, cfg2)
	})
}
