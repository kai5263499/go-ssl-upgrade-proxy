# Go SSL Upgrade Proxy

This is a simple utility to upgrade http requests to https requests. This was developed to solve a problem I had where Nifi refused to pull ElasticSearch records from an endpoint with a self-signed certificate.

Makes heavy use of the [goproxy](https://github.com/elazarl/goproxy) library.

# Build

`GOOS=linux go build -o bin/go-ssl-upgrade-proxy cmd/proxy/main.go`

# Use

`./go-ssl-upgrade-proxy --match 'myes-endpoint:9200'`