FROM golang:alpine as builder

WORKDIR /app
ADD . .

RUN apk update && apk add --no-cache git
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o aug7
RUN chmod -R 700 .

FROM scratch
WORKDIR /apps
COPY --from=builder /apps/aug7 /apps/aug7

EXPOSE 8080
ENTRYPOINT ["/apps/aug7"]
