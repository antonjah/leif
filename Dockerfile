FROM golang:1.13-alpine AS build

WORKDIR /src
COPY . .

ENV GO111MODULE=on
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o leif cmd/leif/leif.go

FROM scratch

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /src/leif /go/bin/leif

ENTRYPOINT ["/go/bin/leif"]