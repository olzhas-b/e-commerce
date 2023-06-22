package clientwrapper

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func SendHttpRequest[Req, Resp any](ctx context.Context, method, url, contentType string, req Req) (resp Resp, err error) {
	bs, err := json.Marshal(req)
	if err != nil {
		return resp, fmt.Errorf("[SendHttpRequest] parsing data: %w", err)
	}

	httpRequest, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(bs))
	if err != nil {
		return resp, fmt.Errorf("[SendHttpRequest] creating request: %w", err)
	}

	httpRequest.Header.Add("Accept", contentType)
	httpRequest.Header.Add("Content-Type", contentType)

	httpResponse, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		return resp, fmt.Errorf("[SendHttpRequest] sending request: %w", err)
	}
	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != http.StatusOK {
		return resp, fmt.Errorf("[SendHttpRequest] statusCode: %d", httpResponse.StatusCode)
	}

	bs, err = io.ReadAll(httpResponse.Body)
	if err != nil {
		return resp, fmt.Errorf("[SendHttpRequest] io.ReadAll: %w", err)
	}

	err = json.Unmarshal(bs, &resp)
	if err != nil {
		return resp, fmt.Errorf("[SendHttpRequest] unmarshal response: %w", err)
	}

	return resp, nil
}
