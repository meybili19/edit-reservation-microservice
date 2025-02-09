FROM golang:1.23.4

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o edit-reservation cmd/main.go

EXPOSE 4001

CMD ["./edit-reservation"]
