
FROM golang:1.24

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o gorm-demo

EXPOSE 8081

CMD ["./gorm-demo"]
