FROM golang:1.21.1-bullseye as builder

WORKDIR /workspace

COPY go.mod go.mod
#COPY go.sum go.sum

RUN go mod download

COPY main.go main.go

RUN go build -a -o example main.go

FROM gcr.io/distroless/base-debian11:latest
WORKDIR /
COPY --from=builder /workspace/example .
ENTRYPOINT ["/example"]
EXPOSE 8080
EXPOSE 8089
