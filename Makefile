BINDIR ?= $(CURDIR)/bin
PATH   := $(BINDIR):$(PATH)

go_vet:
	go vet $$(go list ./...)

go_test: mocks/test.go
	go test -coverprofile=coverage.profile $$(go list ./...)

go_build:
	go build ./...

go_fmt:
	@set -e; \
	GO_FMT=$$(git ls-files '*.go' | grep -v 'vendor/' | xargs gofmt -d -s); \
	if [ -n "$${GO_FMT}" ] ; then \
		echo "Please run gofmt -s"; \
		echo "$$GO_FMT"; \
		exit 1; \
	fi

$(BINDIR)/go-bindata:
	mkdir -p $(BINDIR)
	go build -o $(BINDIR)/go-bindata github.com/jteeuwen/go-bindata/go-bindata

$(BINDIR)/mockgen:
	mkdir -p $(BINDIR)
	go build -o $(BINDIR)/mockgen github.com/golang/mock/mockgen

mocks/test.go: $(BINDIR)/go-bindata $(BINDIR)/mockgen api/interfaces.go
	go generate ./pkg/...
