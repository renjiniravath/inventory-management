FROM golang:alpine as builder

LABEL maintainer="Renji Joseph Sabu <renjiniravath@gmail.com>"

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download 

COPY . .

RUN go build -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

RUN mkdir public

RUN mkdir logs

COPY --from=builder /app/main .
COPY --from=builder /app/.env .
COPY --from=builder /app/public ./public

EXPOSE 8080

CMD ["./main"]