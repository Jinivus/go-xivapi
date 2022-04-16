package xivapi

import (
	"context"
	"fmt"
)

type ServersService service

type ServersResponse []string

func (s *ServersService) Servers(ctx context.Context) (*ServersResponse, *Response, error) {
	result := new(ServersResponse)
	resp, err := s.getServers(ctx, result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

func (s *ServersService) getServers(ctx context.Context, result interface{}) (*Response, error) {

	u := fmt.Sprintf("servers")
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, result)
}
