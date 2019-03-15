package domainr

import (
	"fmt"
	"github.com/google/go-querystring/query"
)

func buildPath(path, query string) (string, error) {
	return fmt.Sprintf("%s?%s", path, query), nil
}

// addOptions adds the parameters in opt as URL query parameters to s.  opt
// must be a struct whose fields may contain "url" tags.
func urlQueryToString(options interface{}) (string, error) {
	v, err := query.Values(options)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}
