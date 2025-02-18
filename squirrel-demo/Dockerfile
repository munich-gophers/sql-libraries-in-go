FROM golang:1.24

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
COPY . .

RUN go build -o sqlx-demo

EXPOSE 8085

CMD ["./sqlx-demo"]
