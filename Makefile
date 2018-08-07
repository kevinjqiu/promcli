.PHONY: build deps proxy

build: deps proxy
	go build

deps:
	dep ensure -v

proxy:
	@ln -s $$(pwd)/_proxy/proxy.go $$(pwd)/vendor/github.com/prometheus/prometheus/promql/proxy.go 2>/dev/null || true

docker:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .
	docker build -t kevinjqiu/promcli .
	docker push kevinjqiu/promcli
