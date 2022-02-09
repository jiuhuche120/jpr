package config

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestConfig_LoadConfig(t *testing.T) {
	config, err := LoadConfig()
	require.Nil(t, err)
	require.Equal(t, "jiuhuche120", config.Owner)
	require.Equal(t, "test_action", config.Repo)
	require.Equal(t, time.Second*10, config.Time)
	require.Equal(t, "master", config.Base)
	require.Equal(t, "release*", config.Head)
}

func TestDefaultConfig(t *testing.T) {
	config := DefaultConfig()
	require.Equal(t, time.Hour*3, config.Time)
	require.Equal(t, "master", config.Base)
	require.Equal(t, "release*", config.Head)
}

func TestPathRoot(t *testing.T) {
	path, err := PathRoot()
	require.Nil(t, err)
	require.Equal(t, "/Users/jiuhuche120/.jpr", path)
}
