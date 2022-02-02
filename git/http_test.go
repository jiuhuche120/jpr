package git

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestNewHttpClient(t *testing.T) {
	client := NewHttpClient()
	resp, err := client.Get("http://www.baidu.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var buffer [5]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
	fmt.Println(result.String())
}
