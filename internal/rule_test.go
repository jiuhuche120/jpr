package internal

import (
	"fmt"
	"testing"

	"github.com/jiuhuche120/jpr/config"
	"github.com/stretchr/testify/require"
)

func TestIsWorkDay(t *testing.T) {
	c, err := config.LoadConfig()
	require.Nil(t, err)
	server := NewServer(c)
	fmt.Println(server.IsWorkingDay())
}
