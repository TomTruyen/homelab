package anilist

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const anilistURL = "https://graphql.anilist.co"

type GraphQLRequest struct {
	Query string `json:"query"`
}

func (s *Service) doGraphQL(
	ctx context.Context,
	query string,
	out any,
) error {
	token, err := s.GetAccessToken(ctx)
	if err != nil {
		return err
	}

	payload := GraphQLRequest{
		Query: query,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		anilistURL,
		bytes.NewBuffer(body),
	)

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf(
			"graphql request failed with: status %d, body: %s",
			resp.StatusCode,
			body,
		)
	}

	return json.NewDecoder(resp.Body).Decode(out)
}
