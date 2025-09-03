FROM golang:1.22.0-alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o jwt_and_rate_limiter ./cmd/main.go

EXPOSE 8080

CMD ["./jwt_and_rate_limiter"]