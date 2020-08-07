package model

import (
	"context"
	"encoding/json"
	"net/http"
)

func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func DecodeGetAllRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}
