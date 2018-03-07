package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/dkozic/goapr/aprclient"
	"github.com/go-kit/kit/endpoint"
)

func makeSearchByRegistryCodeEndpoint(svc AprService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(searchByRegistryCodeRequest)
		result, err := svc.SearchByRegistryCode(req.RegistryCode)
		if err != nil {
			return searchByRegistryCodeResponse{result, err.Error()}, nil
		}
		return searchByRegistryCodeResponse{result, ""}, nil
	}
}

func makeSearchByBusinessNameEndpoint(svc AprService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(searchByBusinessNameRequest)
		result, err := svc.SearchByBusinessName(req.BusinessName)
		if err != nil {
			return searchByBusinessNameResponse{result, err.Error()}, nil
		}
		return searchByBusinessNameResponse{result, ""}, nil
	}
}

func decodeSearchByRegistryCodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request searchByRegistryCodeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeSearchByBusinessNameRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request searchByBusinessNameRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeSearchByRegistryCodeResponse(_ context.Context, r *http.Response) (interface{}, error) {
	var response searchByRegistryCodeResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

func decodeSearchByBusinessNameResponse(_ context.Context, r *http.Response) (interface{}, error) {
	var response searchByBusinessNameResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func encodeRequest(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

type searchByRegistryCodeRequest struct {
	RegistryCode string `json:"registryCode"`
}

type searchByRegistryCodeResponse struct {
	Result aprclient.SearchByRegistryCodeResult `json:"result"`
	Err    string                               `json:"err,omitempty"`
}

type searchByBusinessNameRequest struct {
	BusinessName string `json:"businessName"`
}

type searchByBusinessNameResponse struct {
	Result []aprclient.SearchByBusinessNameResult `json:"result"`
	Err    string                                 `json:"err,omitempty"`
}
