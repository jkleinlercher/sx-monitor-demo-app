FROM golang:1.16-alpine AS build
LABEL org.opencontainers.image.source https://github.com/jkleinlercher/sx-monitor-demo-app
WORKDIR /app
COPY . .
RUN go build -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/main .
EXPOSE 8080
CMD ["./main"]
