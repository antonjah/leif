FROM golang:alpine AS build

WORKDIR /src
COPY . .

ENV GO111MODULE=on
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o gleif cmd/gleif/main.go

FROM scratch

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /src/gleif /go/bin/gleif

ENTRYPOINT ["/go/bin/gleif"]