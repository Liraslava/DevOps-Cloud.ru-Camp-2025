# lirika13/echo-go

FROM golang:1.24-alpine3.21 AS builder
RUN apk update 
WORKDIR /home/build
COPY . .

RUN go build -ldflags="-w -s" -o /home/service/app ./cmd/server

FROM alpine:3.21 AS runner
COPY --from=builder /home/service/app /home/service/app
COPY /run/templates /home/service/templates

WORKDIR /home/service
ENTRYPOINT ["/home/service/app"]

EXPOSE 8000/tcp

