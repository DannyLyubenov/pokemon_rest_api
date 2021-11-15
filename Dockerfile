FROM golang:1.17-stretch as base

WORKDIR /go/src/app

# copy all relevent files over
COPY ./go.sum .
COPY ./go.mod .
COPY ./cmd/ ./cmd/
COPY ./pkg/ ./pkg/

RUN go mod download

# run tests
RUN go test -v ./...

# build the go binary
RUN go build -o /go/bin/app /go/src/app/cmd/

# using the distroless base image
FROM gcr.io/distroless/base as app

# copy the compiled binary
COPY --from=base /go/bin/app /

# set the entrypoint
CMD ["./app"]
