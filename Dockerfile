FROM golang:1.23.1-alpine

WORKDIR /parkirkuy

COPY . .

RUN go mod tidy

RUN go build -o main cmd/app/main.go

CMD ["./main"]