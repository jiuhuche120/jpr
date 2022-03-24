package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfig_LoadConfig(t *testing.T) {
	config, err := LoadConfig()
	require.Nil(t, err)
	fmt.Println(config.DingTalk)
}

func TestPathRoot(t *testing.T) {
	path, err := PathRoot()
	require.Nil(t, err)
	require.Equal(t, "/Users/jiuhuche120/.jpr", path)
}
