package xivapi

import (
	"context"
	"fmt"
)

type SearchService service

type searchParameters struct {
	Query string
}

type Item struct {
	ID      int    `json:"ID"`
	Icon    string `json:"Icon"`
	Name    string `json:"Name"`
	URL     string `json:"Url"`
	URLType string `json:"UrlType"`
	Type    string `json:"_"`
	Score   int    `json:"_Score"`
}

type ItemsSearchResult struct {
	*PaginatedResult
	Items []*Item `json:"Results,omitempty"`
	Speed *int    `json:"SpeedMs,omitempty"`
}

func (s *SearchService) Items(ctx context.Context, query string) (*ItemsSearchResult, *Response, error) {
	result := new(ItemsSearchResult)
	resp, err := s.search(ctx, "Item", query, result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

func (s *SearchService) search(ctx context.Context, searchType string, query string, result interface{}) (*Response, error) {

	u := fmt.Sprintf("search?indexes=%s&string=%s", searchType, query)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, result)
}
