package godalle

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type (
	ImageRequest struct {
	}
	ImageResponse struct {
	}
)

func (c *Client) CreateImageGeneration(
	ctx context.Context,
	request ImageRequest,
) (response ImageResponse, err error) {
	var reqBytes []byte
	reqBytes, err = json.Marshal(request)
	if err != nil {
		return
	}

	urlSuffix := "/images/generations"
	req, err := http.NewRequest("POST", c.fullURL(urlSuffix), bytes.NewBuffer(reqBytes))
	if err != nil {
		return
	}

	req = req.WithContext(ctx)
	err = c.sendRequest(req, &response)
	return
}
