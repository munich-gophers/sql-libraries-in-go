FROM golang:1.24

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
COPY . .

RUN go build -o bun-demo

EXPOSE 8087

CMD ["./bun-demo"]
