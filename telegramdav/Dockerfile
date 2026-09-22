FROM golang:1.24-alpine AS build
WORKDIR /src
COPY . .
WORKDIR /src/telegramdav
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags='-s -w' -o /out/telegramdav ./cmd/telegramdav
FROM alpine:3.22
RUN adduser -D -H app
USER app
COPY --from=build /out/telegramdav /telegramdav
EXPOSE 8080
ENTRYPOINT ["/telegramdav"]
