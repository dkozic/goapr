package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	aprAddress := os.Getenv("APR_URL")
	if aprAddress == "" {
		aprAddress = "http://pretraga2.apr.gov.rs/ObjedinjenePretrage/Search/Search"
	}

	headless := os.Getenv("HEADLESS")
	if headless == "" {
		headless = "true"
	}

	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)

	logger.Log("port", port, "headless", headless, "aprAddress", aprAddress)

	var svc AprService
	svc = aprService{aprAddress, headless == "true"}
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

	logger.Log("msg", "HTTP", "addr", port)
	logger.Log("err", http.ListenAndServe(":"+port, nil))
}
