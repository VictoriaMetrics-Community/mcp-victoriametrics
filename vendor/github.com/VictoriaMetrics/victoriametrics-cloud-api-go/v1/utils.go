package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func requestAPI[R any](ctx context.Context, a *VMCloudAPIClient, method string, body io.Reader, path ...string) (R, error) {
	var result R
	reqURL := a.parsedURL.JoinPath(path...).String()
	req, err := http.NewRequestWithContext(ctx, method, reqURL, body)
	if err != nil {
		return result, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set(AccessTokenHeader, a.apiKey)
	resp, err := a.c.Do(req)
	if err != nil {
		return result, fmt.Errorf("failed to send request: %w", err)
	}
	respBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, fmt.Errorf("failed to read response body: %w", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	if resp.StatusCode/100 != 2 {
		return result, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(respBodyBytes))
	}
	if len(respBodyBytes) > 0 {
		// Special case for string type - just return the response body as a string
		if stringResult, ok := any(&result).(*string); ok {
			*stringResult = string(respBodyBytes)
		} else {
			// For other types, unmarshal as JSON
			if err = json.Unmarshal(respBodyBytes, &result); err != nil {
				return result, fmt.Errorf("failed to unmarshal response body: %w", err)
			}
		}
	}
	return result, nil
}
