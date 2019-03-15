package domainr

import (
	"net/http"
	"net/url"
)

type Search interface {
	Search(query string, opts *SearchOptions) ([]SearchResponse, *http.Response, error)
}
type searchService struct {
	*Client
}

var _ Search = &searchService{}

type SearchOptions struct {
	// Override IP location detection for country-code zones, with a two-character country code
	Location string `url:"location,omitempty"`
	// The domain name of a specific registrar, for filtering results by zones supported by the that registrar
	// Eg. dnsimple.com
	Registrar string `url:"registrar,omitempty"`
	// List of default zones to include in the response (optional).
	Defaults string `url:"defaults,omitempty"`
	// List of keywords, for seeding the results. e.g. a new gTLD like kitchen, or a related keyword like vegan (optional)
	Keywords string `url:"keywords,omitempty"`
}

const (
	searchPath = "search"
)

type Results struct {
	SearchResults []SearchResponse `json:"results"`
}
type SearchResponse struct {
	Domain      string `json:"domain"`
	Host        string `json:"host"`
	Subdomain   string `json:"subdomain"`
	Zone        string `json:"zone"`
	Path        string `json:"path"`
	RegisterURL string `json:"register_url"`
}

func (s *searchService) Search(term string, opts *SearchOptions) ([]SearchResponse, *http.Response, error) {
	q := url.Values{}
	q.Set("mashape-key", s.AuthKey)
	q.Set("query", term)
	query := q.Encode()
	if opts != nil {
		tmp, err := urlQueryToString(opts)
		if err != nil {
			return nil, nil, err
		}
		query += "&" + tmp
	}
	path, err := buildPath(searchPath, query)
	if err != nil {
		return nil, nil, err
	}
	var results Results
	response, err := s.get(path, &results)
	if err != nil {
		return nil, nil, err
	}
	return results.SearchResults, response, nil
}
