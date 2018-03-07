package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

const searchUrl = "http://pretraga2.apr.gov.rs/ObjedinjenePretrage/Search/Search"

func main() {
	var (
		listen = flag.String("listen", ":8080", "HTTP listen address")
	)
	flag.Parse()

	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "listen", *listen, "caller", log.DefaultCaller)

	var svc AprService
	svc = aprService{searchUrl}
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

	logger.Log("msg", "HTTP", "addr", *listen)
	logger.Log("err", http.ListenAndServe(*listen, nil))
}
