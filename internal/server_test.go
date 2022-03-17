package internal

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/jiuhuche120/jpr/config"
	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	c, err := config.LoadConfig()
	require.Nil(t, err)
	server := NewServer(c)
	pulls, err := server.GetAllPullRequests()
	require.Nil(t, err)
	reg := regexp.MustCompile(server.config.Head)
	for i := 0; i < len(pulls); i++ {
		if pulls[i].State == "open" && reg.FindString(pulls[i].Base.Ref) != "" {
			//fmt.Println(pulls[i])
			for j := 0; j < len(pulls); j++ {
				if i == j {
					continue
				}
				if i != j && pulls[i].Title == pulls[j].Title && pulls[j].State == "closed" {
					fmt.Println(pulls[j])
					break
				}
			}
		}
	}
}
