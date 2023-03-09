FROM golang:1.16-alpine AS builder

WORKDIR /app
COPY main.go .

RUN go build -o ./server main.go

FROM alpine

RUN apk --no-cache add curl

WORKDIR /app
COPY --from=builder /app/server ./server

EXPOSE 80

CMD [ "./server" ]