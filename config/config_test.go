package config

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConfig_LoadConfig(t *testing.T) {
	config, err := LoadConfig()
	require.Nil(t, err)
	fmt.Println(config)
}
