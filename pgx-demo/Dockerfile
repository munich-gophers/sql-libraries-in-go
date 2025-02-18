FROM golang:1.24

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o pgx-demo

EXPOSE 8082

CMD ["./pgx-demo"]
