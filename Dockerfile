FROM golang

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/cmd

RUN go build -o main .

WORKDIR /app

EXPOSE 8080

CMD ["./cmd/main"]
