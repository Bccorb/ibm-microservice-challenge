FROM golang:alpine

WORKDIR /app

COPY . .

RUN go mod tidy

EXPOSE 9000

CMD ["go", "run", "main.go", "server"]
