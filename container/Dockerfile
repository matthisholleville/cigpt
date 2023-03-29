FROM golang:1.20.2-alpine3.16 AS builder

ENV CGO_ENABLED=0

WORKDIR /workspace

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN go build -o /workspace/cigpt ./

FROM gcr.io/distroless/static AS production

LABEL org.opencontainers.image.source="https://github.com/matthisholleville/cigpt" \
    org.opencontainers.image.url="https://cigpt.io" \
    org.opencontainers.image.title="cigpt" \
    org.opencontainers.image.vendor="the cigpt-ai maintainers" \
    org.opencontainers.image.licenses="MIT"

WORKDIR /
COPY --from=builder /workspace/cigpt .
USER 65532:65532

ENTRYPOINT ["/cigpt"]