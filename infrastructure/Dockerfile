FROM golang:1.16-alpine

WORKDIR /app

COPY backend/ .

RUN go mod tidy

EXPOSE 8080

CMD ["go", "run", "main.go"]
