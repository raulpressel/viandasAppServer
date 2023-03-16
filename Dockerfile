FROM golang:alpine

WORKDIR /app
COPY . /app

COPY .env /app
COPY cert.pem /app

RUN go build -o main .

CMD ["/app/main"]