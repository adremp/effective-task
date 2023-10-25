package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func JsonUnmarshal[T any](resp *http.Response) (*T, error) {
	defer resp.Body.Close()

	dataB, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var ageResp T
	data := json.Unmarshal(dataB, &ageResp)
	if data != nil {
		return nil, nil
	}

	return &ageResp, nil
}
