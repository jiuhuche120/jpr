package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfig_LoadConfig(t *testing.T) {
	config, err := LoadConfig()
	require.Nil(t, err)
	require.Equal(t, "jiuhuche120", config.Owner)
	require.Equal(t, "XXX", config.Repo)
	require.Equal(t, "0 30 16 * * *", config.Cron)
	require.Equal(t, "master", config.Base)
	require.Equal(t, "release*", config.Head)
}

func TestPathRoot(t *testing.T) {
	path, err := PathRoot()
	require.Nil(t, err)
	require.Equal(t, "/Users/jiuhuche120/.jpr", path)
}
