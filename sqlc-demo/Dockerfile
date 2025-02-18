FROM golang:1.24

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
COPY . .

RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
RUN sqlc generate
RUN go build -o sqlc-demo

EXPOSE 8083

CMD ["./sqlc-demo"]
