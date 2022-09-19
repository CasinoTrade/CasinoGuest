
GIT_COMMIT=`git rev-parse HEAD`
BUILD_DATE=`date +%d-%m-%Y-%H:%M`
VERSION?=0.0.0-develop

all: casino

.PHONY: casino
casino:
	@echo "Version: $(VERSION)"
	@echo "Commit: $(GIT_COMMIT)"
	@echo "Date: $(BUILD_DATE)"

	go get ./...
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w -X main.version=$(VERSION) -X main.commit=$(GIT_COMMIT) -X main.date=$(BUILD_DATE)" -o casinoguest ./cmd/casinoguest/main.go

	@echo "Done"
