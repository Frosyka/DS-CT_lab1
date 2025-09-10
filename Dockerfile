# Stage 1: Сборка приложения
FROM golang:1.23-alpine AS build
WORKDIR /app
COPY src/server.go .
RUN CGO_ENABLED=0 GOOS=linux go build -o server server.go

# Stage 2: Финальный минималистичный образ
FROM gcr.io/distroless/static-debian12
WORKDIR /app
USER nonroot:nonroot
COPY --from=build /app/server /app/server
EXPOSE 8080
ENV PORT=8080
HEALTHCHECK --interval=30s --timeout=3s CMD ["/bin/wget", "--quiet", "--tries=1", "--spider", "http://127.0.0.1:${PORT}/health"]
ENTRYPOINT ["/app/server"]