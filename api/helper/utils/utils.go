package utils

import (
	"net/http"
	"strconv"
	"strings"
	"tender/api/tokens"
	"tender/config"
)

type QueryParams struct {
	Page  int64
	Limit int64
}

func ParseQueryParams(queryParams map[string][]string) (*QueryParams, []string) {
	params := QueryParams{
		Page:  1,
		Limit: 10,
	}
	var errStr []string
	var err error

	for key, value := range queryParams {
		if key == "page" {
			params.Page, err = strconv.ParseInt(value[0], 10, 64)
			if err != nil {
				errStr = append(errStr, "Invalid `page` param")
			}
			continue
		}

		if key == "limit" {
			params.Limit, err = strconv.ParseInt(value[0], 10, 64)
			if err != nil {
				errStr = append(errStr, "Invalid `limit` param")
			}
			continue
		}
	}

	return &params, errStr
}

func GetClaimsFromToken(request *http.Request, cfg *config.Config) (map[string]interface{}, error) {
	token := request.Header.Get("Authorization")

	if token == "" {
		return map[string]interface{}{
			"role": "unauthorized",
			"sub":  nil,
			"exp":  nil,
			"iat":  nil,
		}, nil
	}

	if strings.Contains(token, "Bearer ") {
		token = strings.Split(token, "Bearer ")[1]
	}

	claims, err := tokens.ExtractClaim(token, []byte(cfg.SigningKey))
	if err != nil {
		return nil, err
	}

	return claims, nil
}
