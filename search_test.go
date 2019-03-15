package domainr

import (
	"fmt"
	"os"
	"testing"
)

func Test_Search(t *testing.T) {
	client := NewClient(AuthKey(os.Getenv("AUTH_KEY")))
	so := &SearchOptions{
		Registrar: "dnsimple.com",
		Defaults:  "com,co,net,us",
	}
	search, _, err := client.Search.Search("randy's accounting", so)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	for _, item := range search {
		fmt.Println(item.Domain)
		fmt.Println(item.Zone)
	}
}
