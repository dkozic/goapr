package main

import (
	"time"

	"github.com/dkozic/goapr/aprclient"
	"github.com/go-kit/kit/log"
)

func loggingMiddleware(logger log.Logger) ServiceMiddleware {
	return func(next AprService) AprService {
		return logmw{logger, next}
	}
}

type logmw struct {
	logger log.Logger
	next   AprService
}

func (mw logmw) SearchByRegistryCode(input string) (output aprclient.SearchByRegistryCodeResult, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "SearchByRegistryCode",
			"input", input,
			"output", output.MainData.BusinessName,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.SearchByRegistryCode(input)
	return
}

func (mw logmw) SearchByBusinessName(input string) (output []aprclient.SearchByBusinessNameResult, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "SearchByBusinessName",
			"input", input,
			"output", len(output),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.SearchByBusinessName(input)
	return
}
