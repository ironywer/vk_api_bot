FROM golang:1.21

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o bot ./cmd

EXPOSE 8080

CMD ["./bot"]
