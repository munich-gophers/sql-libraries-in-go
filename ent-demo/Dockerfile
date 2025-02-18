FROM golang:1.24

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
COPY . .

RUN go build -o ent-demo

EXPOSE 8086

CMD ["./ent-demo"]
