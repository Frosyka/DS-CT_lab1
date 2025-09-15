FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY src/server.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server server.go

FROM alpine:3.20
RUN adduser -D -u 10001 nonroot
WORKDIR /app
COPY --from=builder /app/server .
RUN chown nonroot:nonroot /app/server
USER nonroot
ENV PORT=8092
EXPOSE $PORT
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:$PORT/health || exit 1
CMD ["./server"]