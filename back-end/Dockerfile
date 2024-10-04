FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

RUN apk --no-cache add ca-certificates openssl

EXPOSE 8000

CMD ["sh", "-c", "./script.sh && ./main"]
