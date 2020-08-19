package model

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dhuki/rest-template/pkg/testing/domain/entity"
)

func DecodeCreateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request entity.TestTable
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}
	return request, nil
}
