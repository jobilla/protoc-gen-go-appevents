FROM golang:latest as build

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

ENV CGO_ENABLED=0
RUN go build -ldflags="-s -w" -o protoc-gen-go-appevents

FROM scratch

COPY --from=build /build/protoc-gen-go-appevents /protoc-gen-go-appevents
