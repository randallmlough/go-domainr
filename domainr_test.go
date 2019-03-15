package domainr

import (
	"os"
)

var dmnr *Client

func init() {
	dmnr = NewClient(AuthKey(os.Getenv("AUTH_KEY")))
}
