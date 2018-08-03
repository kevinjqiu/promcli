build: proxy
	go build

proxy:
	@ln -s $$(pwd)/_proxy/proxy.go $$(pwd)/vendor/github.com/prometheus/prometheus/promql/proxy.go 2>/dev/null || true
