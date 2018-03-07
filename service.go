package main

import (
	"errors"

	"github.com/dkozic/goapr/aprclient"
)

// AprService provides operations on APR.
type AprService interface {
	SearchByRegistryCode(string) (aprclient.SearchByRegistryCodeResult, error)
	SearchByBusinessName(string) ([]aprclient.SearchByBusinessNameResult, error)
}

type aprService struct {
	url string
}

func (svc aprService) SearchByRegistryCode(registryCode string) (aprclient.SearchByRegistryCodeResult, error) {
	var res aprclient.SearchByRegistryCodeResult
	if registryCode == "" {
		return res, ErrEmpty
	}
	client := aprclient.New(svc.url)
	res, err := client.SearchByRegistryCode(registryCode)
	return res, err
}

func (svc aprService) SearchByBusinessName(businessName string) ([]aprclient.SearchByBusinessNameResult, error) {
	if businessName == "" {
		return nil, ErrEmpty
	}
	client := aprclient.New(svc.url)
	if res, err := client.SearchByBusinessName(businessName); err != nil {
		return nil, err
	} else {
		return res, nil
	}

}

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty input")

// ServiceMiddleware is a chainable behavior modifier for AprService.
type ServiceMiddleware func(AprService) AprService
