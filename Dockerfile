FROM golang:1.10.3

# Download and install the latest release of dep
ADD https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

# copy project
WORKDIR $GOPATH/src/github.com/tusupov/gofetchtask
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only
COPY . ./

# run test
RUN go test -v ./...
RUN go test -bench=. -v ./...
RUN rm -rf *_test.go

# build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gofetchtask .

FROM alpine:latest

# add certificates for https connections
RUN apk --no-cache add ca-certificates

# copy
WORKDIR /app/
COPY --from=0 /go/src/github.com/tusupov/gofetchtask/gofetchtask .

EXPOSE $PORT

CMD ["./gofetchtask", "p", "$PORT"]
