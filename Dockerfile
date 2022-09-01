FROM golang:alpine

WORKDIR /app
COPY . /app

COPY b.jpg /var/www/default/htdocs/public/


RUN go build -o main .

CMD ["/app/main"]