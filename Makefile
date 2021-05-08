UNFORMATTED := $(shell gofmt -l . )
BUILD_SERVER := $(shell mkdir -p bin && cd server/ && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ../bin/pgmeta -v .)
ci:
	[ -z ${UNFORMATTED} ] && exit 0
	find . -name go.mod -execdir go test ./... \;

build-server:
	exec ${BUILD_SERVER}

deploy:
	exec ${BUILD_SERVER} && bin/pgmeta
