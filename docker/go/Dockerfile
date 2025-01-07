FROM golang:1.23.3

ENV GO111MODULE on
WORKDIR /go-clean-app
RUN apt-get update && apt-get install -y vim postgresql-client
RUN go install github.com/cosmtrek/air@v1.27.4 && \
  go install github.com/pressly/goose/v3/cmd/goose@latest && \
  go install github.com/go-delve/delve/cmd/dlv@latest && \
  go install golang.org/x/tools/cmd/goimports@latest && \
  go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

CMD ["air", "-c", ".air.toml"]