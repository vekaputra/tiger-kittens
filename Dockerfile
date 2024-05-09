# Stage build: will build the go into single binary
FROM golang:1.21.10-bookworm AS build
WORKDIR /go/src/app
COPY ./cmd ./cmd
COPY ./db ./db
COPY ./env ./env
COPY ./internal ./internal
COPY ./main.go ./
COPY ./go.mod ./
COPY ./go.sum ./

RUN go env -w CGO_ENABLED=0
RUN go env -w GO111MODULE="on"
RUN go env -w GOOS=linux
RUN go env -w GOARCH=amd64
RUN git config --global --add url."git@github.com:".insteadOf "https://github.com/"

# Install dependencies using docker cache.
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

# Build binary using docker cache.
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -v .

CMD ["/go/src/app/tiger-kittens", "serve"]