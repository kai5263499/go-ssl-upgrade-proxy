package main

import (
	"flag"
	"fmt"
	"net/http"
	"regexp"

	"github.com/Sirupsen/logrus"
	"github.com/elazarl/goproxy"
)

var (
	port  int
	match string
)

func main() {

	flag.IntVar(&port, "port", 18088, "proxy port")
	flag.StringVar(&match, "match", "", "urls matching this regex pattern will be")

	flag.Parse()

	logrus.SetLevel(logrus.DebugLevel)

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true

	urlMatcher := regexp.MustCompile(match)

	proxy.OnRequest(goproxy.UrlMatches(urlMatcher)).DoFunc(
		func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			r.URL.Scheme = "https"
			return r, nil
		})

	logrus.Debugf("listening on %d", port)
	logrus.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), proxy))
}
