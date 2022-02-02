package git

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClient_Get(t *testing.T) {
	client := NewClient(
		AddHeader("Accept", "application/vnd.github.v3+json"),
		AddHeader("state", "open"),
	)
	res := client.Get("https://api.github.com/repos/meshplus/bitxhub/pulls")
	var pulls []PullRequest
	err := json.Unmarshal([]byte(res), &pulls)
	require.Nil(t, err)
	fmt.Println(pulls[0].Head.Ref)
}
