build:
	go build

proxy:
	ln -s $(pwd)/_proxy/proxy.go $(pwd)/vendor/github.com/prometheus/prometheus/promql/proxy.go
