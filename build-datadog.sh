#!/bin/bash

go install golang.org/x/vuln/cmd/govulncheck@latest
go install golang.org/x/tools/cmd/deadcode@latest
go install github.com/mgechev/revive@latest

go install github.com/DataDog/orchestrion@v1.6.1

echo "adding orchestrion pin"
orchestrion pin

gofmt -s -w .

revive ./...

gocyclo -over 15 .

go mod tidy

govulncheck ./...

deadcode ./cmd/*

go env -w CGO_ENABLED=1

go test -race ./...

go env -w CGO_ENABLED=0

orchestrion go build -o ~/go/bin/apiping-datadog ./cmd/apiping

go env -u CGO_ENABLED
