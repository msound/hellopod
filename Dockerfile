FROM docker.io/golang:1.24 AS build-stage

ENV GOOS=linux GOARCH=amd64 CGO_ENABLED=0

WORKDIR /app

COPY go.* .

RUN go mod download

COPY . .

RUN go build -o hellopod

######################

FROM alpine:3

RUN apk update \
  && apk add --no-cache --purge ca-certificates tzdata \
  && rm -rf /var/cache/apk/* /tmp/*

RUN adduser -D appuser --uid 3000

COPY --from=build-stage /app/hellopod /app/hellopod

USER appuser

CMD ["/app/hellopod"]
