package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

const APR_URL = "http://pretraga2.apr.gov.rs/ObjedinjenePretrage/Search/Search"

func main() {
	listen := os.Getenv("PORT")
	if listen == "" {
		listen = "8080"
	}

	aprUrl := os.Getenv("APR_URL")
	if aprUrl == "" {
		aprUrl = APR_URL
	}

	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "listen", listen, "caller", log.DefaultCaller)

	var svc AprService
	svc = aprService{aprUrl}
	svc = loggingMiddleware(logger)(svc)

	searchByRegistryCodeHandler := httptransport.NewServer(
		makeSearchByRegistryCodeEndpoint(svc),
		decodeSearchByRegistryCodeRequest,
		encodeResponse,
	)
	searchByBusinessNameHandler := httptransport.NewServer(
		makeSearchByBusinessNameEndpoint(svc),
		decodeSearchByBusinessNameRequest,
		encodeResponse,
	)

	http.Handle("/searchByRegistryCode", searchByRegistryCodeHandler)
	http.Handle("/searchByBusinessName", searchByBusinessNameHandler)

	logger.Log("msg", "HTTP", "addr", listen)
	logger.Log("err", http.ListenAndServe(":"+listen, nil))
}
